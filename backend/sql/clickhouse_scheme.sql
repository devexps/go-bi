CREATE OR REPLACE TABLE realtime_warehousing
(
    event_name  String,
    report_data String,
    create_time DateTime('Asia/Ho_Chi_Minh') DEFAULT now()
)
    ENGINE = MergeTree()
    PARTITION BY (toYYYYMMDD(create_time))
        ORDER BY (toYYYYMMDD(create_time), event_name)
        TTL create_time + toIntervalMonth(3)
        SETTINGS index_granularity = 8192;