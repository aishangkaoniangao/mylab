function generate_bash_file()
{
    src=$1
    dest=${src}'.sh'
    echo $src

    #clear file data
    echo '' > $dest

    cat ${src} >> $dest
    cat ${RUN_TPL} >> $dest
}

ROOT=`pwd`
meta_dir=$ROOT'/conf/*'
RUN_TPL=$ROOT'/conf/run.tpl'
for file in $meta_dir
do
    if [ ${file:0-5} = ".meta" ]
    then
        meta_name=$file
        generate_bash_file $meta_name
    fi
done
