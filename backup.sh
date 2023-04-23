# 需要备份的仓库, 用户名/仓库名
REPOSITORIES=("TheAlgorithms/Go" "avelino/awesome-go" "go-gorm/gorm" "unknwon/the-way-to-go_ZH_CN" "Dreamacro/clash" "v2fly/v2ray-core")

# 因为最终保存的是压缩文件
# 所以在临时目录进行操作
TMP_DIR=$(mktemp -d)
cd $TMP_DIR

# 遍历仓库备份
for repository in ${REPOSITORIES[@]}; do
  # 获取仓库名
  tmp_array=(${repository//\// })
  repository_name=${tmp_array[1]}

  # clone 仓库到本地
  git clone https://github.com/${repository}

  data_time=$(date "+-%Y-%m-%d-%H-%M-%S")
  # 对仓库进行压缩
  zip_name=${repository_name}${data_time}.zip
  echo ${zip_name}
  zip -r $zip_name $repository_name

  # 把压缩文件移到备份目录
  # -f 强行覆盖之前的备份
  mv -f $zip_name ~/github_backup/

  # 休眠等待下一个备份, 五分钟
  sleep 10
done