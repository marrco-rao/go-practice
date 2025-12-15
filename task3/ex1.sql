create table students
(
    id    int auto_increment
        primary key,
    name  varchar(50) null,
    age   int         null,
    grade varchar(30) null
);

insert into test1.students(name, age, grade) value ('张三',20,'三年级');
insert into test1.students(name, age, grade) value ('李四',19,'三年级');
insert into test1.students(name, age, grade) value ('王五',16,'二年级');
insert into test1.students(name, age, grade) value ('赵六',17,'二年级');
insert into test1.students(name, age, grade) value ('田七',22,'四年级');

insert into test1.students(name, age, grade) value ('阿斯顿',15,'二年级');
insert into test1.students(name, age, grade) value ('阿斯顿2',14,'二年级');
insert into test1.students(name, age, grade) value ('阿斯顿3',13,'二年级');
insert into test1.students(name, age, grade) value ('阿斯顿4',12,'二年级');


select * from students where age>18;

update students set grade='四年级' where name='张三';

delete from students where age<15;



-----------------编写内容：控制台无法执行，需套入存储过程--
create table accounts
(
    id           int auto_increment
        primary key,
    balance      int         null,
    account_name varchar(30) null
);

-- 没有用外键
create table transactions
(
    id              int auto_increment
        primary key,
    from_account_id int null,
    to_account_id   int null,
    amount          int null
);


set @amount = 100;
set @from_account_id = 1 ; -- 账户A
set @to_account_id = 2;   -- 账户B

start transaction;
select balance into @from_balance
from accounts
where id = @from_account_id
    for update;

if @from_balance < @amount then
    rollback;
select '转账失败：余额不足' as message;
else
update accounts set balance = balance - @amount where id = @from_account_id;
update accounts set balance = balance + @amount where id = @to_account_id;

insert into transactions(from_account_id, to_account_id, amount) VALUE (@from_account_id,@to_account_id,@amount);

commit ;
select '转账成功' as message;
end if;

-- 参考AI 存储过程代码------
DELIMITER $$

CREATE PROCEDURE transfer_money(
    IN p_from INT,
    IN p_to INT,
    IN p_amount DECIMAL(10,2)
)
BEGIN
    DECLARE v_balance DECIMAL(10,2);

    -- 开始事务
START TRANSACTION;

-- 1. 查询 A 的余额并锁定
SELECT balance INTO v_balance
FROM accounts
WHERE id = p_from
    FOR UPDATE;

-- 2. 判断余额是否足够
IF v_balance < p_amount THEN
        -- 回滚
        ROLLBACK;
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Insufficient balance';
END IF;

    -- 3. 扣款
UPDATE accounts
SET balance = balance - p_amount
WHERE id = p_from;

-- 4. 加钱
UPDATE accounts
SET balance = balance + p_amount
WHERE id = p_to;

-- 5. 记录交易
INSERT INTO transactions(from_account_id, to_account_id, amount)
VALUES (p_from, p_to, p_amount);

-- 提交事务
COMMIT;
END$$

DELIMITER ;

--  调用
CALL transfer_money(1, 2, 100);
---------------
