/*
Create a stream to store transactions for customers
*/
CREATE STREAM customer_transactions (
    tx_id VARCHAR KEY,
    customer_id INT,
    timestamp VARCHAR,
    amount DECIMAL(12, 2)
) WITH (
    kafka_topic = 'customer_transactions',
    partitions = 8,
    value_format = 'json',
    timestamp = 'timestamp',
    timestamp_format = 'yyyy-MM-dd''T''HH:mm:ss'
);

/*
Create a stream to store anomalies in customer transactions
*/
CREATE TABLE possible_customer_anomalies WITH (
    kafka_topic = 'possible_customer_anomalies',
        VALUE_FORMAT='JSON',
)   AS
    SELECT customer_id AS `customer_id_key`,
           as_value(customer_id) AS `customer_id`,
           count(*) AS `n_attempts`,
           sum(amount) AS `total_amount`,
           collect_list(tx_id) AS `tx_ids`,
           WINDOWSTART as `start_boundary`,
           WINDOWEND as `end_boundary`
    FROM customer_transactions
    WINDOW TUMBLING (SIZE 30 SECONDS, RETENTION 1000 DAYS)
    GROUP BY customer_id
    HAVING count(*) >= 3
    EMIT CHANGES;


/*
Initiate a customer transactions event (change the values accordingly)
*/
INSERT INTO customer_transactions (
    customer_id, tx_id, timestamp, amount
) VALUES (
    2,
    '358579699210099',
    '2020-04-22T03:19:58',
    50.25
);