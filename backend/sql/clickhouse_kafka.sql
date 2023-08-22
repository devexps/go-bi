-- realtime_warehousing kafka
CREATE OR REPLACE TABLE realtime_warehousing_kafka
(
    createTime DateTime('Asia/Ho_Chi_Minh'),
    eventName  String,
    reportData String
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = '127.0.0.1:9092',
            kafka_topic_list = 'logger.report.event',
            kafka_group_name = 'ck-saver',
            kafka_format = 'JSONEachRow',
            kafka_skip_broken_messages = 1;

-- realtime_warehousing from kafka
CREATE MATERIALIZED VIEW realtime_warehousing_mv TO realtime_warehousing AS
SELECT createTime as create_time,
       eventName as event_name,
       reportData as report_data
FROM realtime_warehousing_kafka;