/*
Create a stream to store transactions for users
*/
CREATE STREAM user_transactions (
    tx_id VARCHAR KEY,
    user_id INT,
    timestamp VARCHAR,
    amount DECIMAL(12, 2)
) WITH (
    kafka_topic = 'user_transactions',
    partitions = 8,
    value_format = 'json',
    timestamp = 'timestamp',
    timestamp_format = 'yyyy-MM-dd''T''HH:mm:ss'
);

/*
Create a stream to store anomalies in user transactions
*/
CREATE TABLE possible_user_anomalies WITH (
    kafka_topic = 'possible_user_anomalies',
        VALUE_FORMAT='JSON',
)   AS
    SELECT user_id AS `user_id_key`,
           as_value(user_id) AS `user_id`,
           count(*) AS `n_attempts`,
           sum(amount) AS `total_amount`,
           collect_list(tx_id) AS `tx_ids`,
           WINDOWSTART as `start_boundary`,
           WINDOWEND as `end_boundary`
    FROM user_transactions
    WINDOW TUMBLING (SIZE 30 SECONDS, RETENTION 1000 DAYS)
    GROUP BY user_id
    HAVING count(*) >= 3
    EMIT CHANGES;


/*
Initiate a user transactions event
*/
INSERT INTO user_transactions (
    user_id, tx_id, timestamp, amount
) VALUES (
    2,
    '358579699210099',
    '2020-04-22T03:19:58',
    50.25
);