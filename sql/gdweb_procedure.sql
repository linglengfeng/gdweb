-- -------------------------------------------------------------------------------------------------------------------
DROP PROCEDURE IF EXISTS `sp_user_insert`;
delimiter ;;
CREATE PROCEDURE `sp_user_insert`(in in_account varchar(128))
BEGIN
select count(*) into @count from `user` where `account` = in_account;
if @count = 0 then
    insert into `user` (`account`) values (in_account);
    select id from `user` where `account` = in_account;
else
    select id from `user` where `account` = in_account;
end if;
END
;;
delimiter ;

