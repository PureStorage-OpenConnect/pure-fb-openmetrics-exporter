import re
from prometheus_client.core import GaugeMetricFamily


class ArrayHardwareMetrics():
    """
    Base class for FlashBlade Prometheus hardware metrics
    """
    def __init__(self, fb_client):
        self.array_hardware_status = None
        self.hardware = fb_client.hardware()

    def _array_hardware_status(self):
        """
        Create metric of gauge types for components status.

        """
        self.array_hardware_status = GaugeMetricFamily(
                                  'purefb_hardware_health',
                                  'FlashBlade hardware component health status',
                                   labels=['chassis', 'fabric_module', 'type',
                                           'name', 'index', 'slot'])
        re_ch = re.compile(r"^(CH\d+)$")
        re_fb = re.compile(r"^(CH\d+)\.(FB[0-9]+)$")
        re_fm = re.compile(r"^(CH\d+)\.(FM[0-9]+)$")
        re_eth = re.compile(r"^(CH\d+)\.(FM\d+)\.ETH([0-9]+)(\.[0-9]+)?$")
        re_fan = re.compile(r"^(CH\d+)\.(FM\d+)\.FAN([0-9]+)$")
        re_pwr = re.compile(r"^(CH\d+)\.(PWR[0-9]+)$")

        for comp in self.hardware:
            if comp.status == 'not_installed':
                continue
            c_name = comp.name
            c_index = comp.index or ''
            c_slot = comp.slot or ''
            chassis = ''
            fabric_module = ''
            
            if comp.status == 'healthy':
                c_state = 1
            elif comp.status == 'unused':
                c_state = 2
            else:
                c_state = 0

            c_type = comp.type
            if c_type == 'ch':               # Chassis
                detail = re_ch.match(c_name)
                chassis = detail.group(1)
            elif c_type == 'fb':             # Flash Blade
                detail = re_fb.match(c_name)
                chassis = detail.group(1)
            elif c_type == 'fm':             # Fabric Module
                detail = re_fm.match(c_name)
                chassis = detail.group(1)
                fabric_module = detail.group(2)
            elif c_type == 'eth':            # Ethernet card
                detail = re_eth.match(c_name)
                chassis = detail.group(1)
                fabric_module = detail.group(2)
            elif c_type == 'fan':            # Fan
                detail = re_fan.match(c_name)
                chassis = detail.group(1)
                fabric_module = detail.group(2)
            elif c_type == 'pwr':            # Power
                detail = re_pwr.match(c_name)
                chassis = detail.group(1)

            self.array_hardware_status.add_metric([chassis, 
                                                   fabric_module,
                                                   c_type,
                                                   c_name,
                                                   str(c_index),
                                                   str(c_slot)], c_state)

    def get_metrics(self):
        self._array_hardware_status()
        yield self.array_hardware_status
