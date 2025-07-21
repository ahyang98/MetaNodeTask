insert into students  (name, age, grade) values ('张三', 28, '三年级');
select * from students where age > 18;
update students set name  = '张三' where grade = '三年级';
delete from students where age < 15;

DELIMITER $$
CREATE PROCEDURE safe_transfer()
BEGIN
    DECLARE balance_check INT;
    START TRANSACTION;
    SELECT accounts.balance FROM accounts WHERE ID=1 FOR UPDATE ;
    SELECT accounts.balance FROM accounts WHERE ID=2 FOR UPDATE ;
    UPDATE accounts SET balance = balance - 100 where id = 1;
    UPDATE accounts SET balance = balance + 100 where id = 2;
    insert into transactions (from_account_id, to_account_id, amount) VALUES (0,1,100);
    SELECT balance INTO balance_check FROM accounts WHERE id = 1 FOR UPDATE;
    IF balance_check < 100 THEN
        ROLLBACK;
    ELSE
       commit;
    END IF;
END $$
DELIMITER ;
