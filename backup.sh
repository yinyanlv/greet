#!/usr/bin/env bash

#获取当前时间
date_now=$(date "+%Y%m%d-%H%M%S")
backup_folder=/root/backup/prot-backup/data
db_name="prot"
mysql_container_name="mysql"
#定义备份文件名
filename="${db_name}_${date_now}.sql"
backup_filename="${backup_folder}/${filename}"
echo "Starting backup ${db_name} at ${date_now}."
#/usr/bin/mysqldump -u${username} -p${password}  --lock-all-tables --flush-logs ${db_name} > ${backup_filename}
docker exec -i ${mysql_container_name} bash << 'EOF'
mysqldump prot -uroot -p111111 > /prot.sql
exit
EOF
rm -rf ${backup_folder}
mkdir ${backup_folder}
docker cp ${mysql_container_name}:/prot.sql ${backup_folder}
cd ${backup_folder}
tar zcvf ${filename}.tar.gz prot.sql 
rm prot.sql

date_end=$(date "+%Y%m%d-%H%M%S")
echo "Finish backup ${db_name} at ${date_end}."

git add -A
git commit -m "add: backup ${filename}"
git push

#crontab -e
#0 0 * * * bash /root/backup/prot-backup/backup.sh
#sudo service cron restart
