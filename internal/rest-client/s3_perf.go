package restclient


type S3Performance struct {
    Name                  string  `json:"name"`
    Id                    string  `json:"id"`
    OthersPerSec          float64 `json:"others_per_sec"`
    ReadBucketsPerSec     float64 `json:"read_buckets_per_sec"`
    ReadObjectsPerSec     float64 `json:"read_objects_per_sec"`
    WriteBucketsPerSec    float64 `json:"write_buckets_per_sec"`
    WriteObjectsPerSec    float64 `json:"write_objects_per_sec"`
    Time                  int     `json:"time"`
    UsecPerOtherOp        float64 `json:"usec_per_other_op"`
    UsecPerReadBucketOp   float64 `json:"usec_per_read_bucket_op"`
    UsecPerReadObjectOp   float64 `json:"usec_per_read_object_op"`
    UsecPerWriteBucketOp  float64 `json:"usec_per_write_bucket_op"`
    UsecPerWriteObjectOp  float64 `json:"usec_per_write_object_op"`
}
