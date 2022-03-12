
def test_array(fb_client):
    arr1 = fb_client.arrays()
    arr2 = fb_client.arrays()
    assert len(arr1) > 0
    assert arr1 == arr2

def test_hardware(fb_client):
    hw1 = fb_client.hardware()
    hw2 = fb_client.hardware()
    assert len(hw1) > 0
    assert hw1 == hw2

def test_alerts(fb_client):
    alert1 = fb_client.alerts()
    alert2 = fb_client.alerts()
    assert alert1 == alert2

#def test_arrays_clients_performance(fb_client):
#    cli_perf1 = fb_client.arrays_clients_performance()
#    cli_perf2 = fb_client.arrays_clients_performance()
#    assert cli_perf1 == cli_perf2

def test_arrays_performance(fb_client):
    array_perf1 = fb_client.arrays_performance()
    array_perf2 = fb_client.arrays_performance()
    assert len(array_perf1) > 0
    assert array_perf1 == array_perf2

def test_arrays_http_specific_performance(fb_client):
    array_perf1 = fb_client.arrays_http_specific_performance()
    array_perf2 = fb_client.arrays_http_specific_performance()
    assert len(array_perf1) > 0
    assert array_perf1 == array_perf2

def test_arrays_nfs_specific_performance(fb_client):
    array_perf1 = fb_client.arrays_nfs_specific_performance()
    array_perf2 = fb_client.arrays_nfs_specific_performance()
    assert len(array_perf1) > 0
    assert array_perf1 == array_perf2

def test_arrays_s3_specific_performance(fb_client):
    array_perf1 = fb_client.arrays_s3_specific_performance()
    array_perf2 = fb_client.arrays_s3_specific_performance()
    assert len(array_perf1) > 0
    assert array_perf1 == array_perf2

def test_arrays_space(fb_client):
    array_space1 = fb_client.arrays_space()
    array_space2 = fb_client.arrays_space()
    assert len(array_space1) > 0
    assert array_space1 == array_space2

def test_buckets(fb_client):
    buckets1 = fb_client.buckets()
    buckets2 = fb_client.buckets()
    assert buckets1 == buckets2

def test_buckets_performance(fb_client):
    buckets_perf1 = fb_client.buckets_performance()
    buckets_perf2 = fb_client.buckets_performance()
    buckets1 = fb_client.buckets()
    assert buckets_perf1 == buckets_perf2
    assert (len(buckets1) == 0 and len(buckets_perf1) == 0) or (len(buckets1) > 0 and len(buckets_perf1) > 0)

def test_buckets_s3_specific_performance(fb_client):
    buckets_perf1 = fb_client.buckets_s3_specific_performance()
    buckets_perf2 = fb_client.buckets_s3_specific_performance()
    buckets1 = fb_client.buckets()
    assert buckets_perf1 == buckets_perf2
    assert (len(buckets1) == 0 and len(buckets_perf1) == 0) or (len(buckets1) > 0 and len(buckets_perf1) > 0)

def test_bucket_replica_links(fb_client):
    buckets_rlink1 = fb_client.bucket_replica_links()
    buckets_rlink2 = fb_client.bucket_replica_links()
    assert buckets_rlink1 == buckets_rlink2

def test_file_systems(fb_client):
    file_system1 = fb_client.file_systems()
    file_system2 = fb_client.file_systems()
    assert file_system1 == file_system2

def test_file_systems_performance(fb_client):
    file_system_perf1 = fb_client.file_systems_performance()
    file_system_perf2 = fb_client.file_systems_performance()
    file_system1 = fb_client.file_systems()
    assert file_system_perf1 == file_system_perf2
    assert (len(file_system1) == 0 and len(file_system_perf1) == 0) or (len(file_system1) > 0 and len(file_system_perf1) > 0)

def test_file_system_replica_links(fb_client):
    file_system_rlink1 = fb_client.file_system_replica_links()
    file_system_rlink2 = fb_client.file_system_replica_links()
    assert file_system_rlink1 == file_system_rlink2

def test_hardware_connectors_performance(fb_client):
    hw_nic1 = fb_client.hardware_connectors_performance()
    hw_nic2 = fb_client.hardware_connectors_performance()
    assert len(hw_nic1) > 0
    assert hw_nic1 == hw_nic2

def test_usage_groups(fb_client):
    usage1 = fb_client.usage_groups()
    usage2 = fb_client.usage_groups()
    assert usage1 == usage2

def test_usage_users(fb_client):
    usage1 = fb_client.usage_users()
    usage2 = fb_client.usage_users()
    assert usage1 == usage2
