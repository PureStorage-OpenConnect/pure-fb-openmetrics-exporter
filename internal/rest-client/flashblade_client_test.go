package restclient

import (
    "testing"
    "fmt"
)

func TestNew(t *testing.T) {
//    New("10.239.112.80", "T-973ff524-23f7-465b-8320-5eaad9d6a738")
    fb := New("10.225.112.74", "T-da03a759-bb3f-4995-9c6c-fee48d7c1a98")
//    fb := New("10.11.1.1", "T-da03a759-bb3f-4995-9c6c-fee48d7c1a91")
//    arr := fb.GetArrays()
//    fmt.Println(arr)
//    alrt := fb.GetAlerts("state='open'")
//    fmt.Println(alrt)
//    fmt.Println(fb.Error)
//    perf := fb.GetClientsPerformance()
//    perf := fb.GetArraysPerformance()
//    perf := fb.GetArraysHttpPerformance()
//    perf := fb.GetArraysNfsPerformance()
//    perf := fb.GetArraysS3Performance()
//    perf := fb.GetBucketsPerformance()
//    b := fb.GetBuckets()
//    perf := fb.GetBucketsPerformance(b)
//    perf := fb.GetBucketsS3Performance(b)
//    space := fb.GetArraysSpace()
//    fmt.Println(space)
//    f := fb.GetFileSystems()
//    perf := fb.GetFileSystemsPerformance(f, "NFS")
//    perf := fb.GetArraysPerformanceReplication()
//    h := fb.GetHardware()
//    perf := fb.GetUsageGroups(f)
//    bl := fb.GetBlades()
//    fmt.Println(bl)
    perf := fb.GetHwConnectorsPerformance()
    fmt.Println(perf)
    fmt.Println(fb.Error)
}
