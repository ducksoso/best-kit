#!/bin/bash

# set -x

if [ $# -ne 1 ]; then
    echo "Usage: $0 packages_dir"
    exit 1
fi

[ ! -d $1 ] && echo "Error: you should provide a directory." && exit 1

dest=$1
dest=${dest%/}

if ! echo $dest |grep -q "^/"; then
    echo "Error: please use the absolute path."
    exit 1
fi

if ! ls $dest | egrep -q "(gz|zip)$"; then
   echo "Note: nothing need to do."
   exit 0
fi

#---------------------------------------------
TOPDIR=$(cd $(dirname "$0") && pwd)
tmpdir=`mktemp -d`
#---------------------------------------------

for i in `ls ${dest}/{*.gz,*.zip} 2>/dev/null`
do
    rm -rf $tmpdir/*
    cp $i $tmpdir
    cd $tmpdir
    package_arch_name=`ls`
    if echo $package_arch_name | grep -q "gz$"; then
        tar xf $package_arch_name
        gz_suffix=1
    else
        unzip $package_arch_name
        gz_suffix=0
    fi
    rm -rf $package_arch_name
    package_name=`ls`
    cd $package_name
    if ls |grep -q "egg-info"; then
        python setup.py egg_info
        python setup.py build
        cd ..
        if [ $gz_suffix -eq 1 ]; then
            tar czf $package_arch_name $package_name
        else
            zip -r $package_arch_name $package_name
        fi
        rm -rf $i
        cp $package_arch_name $dest/
    fi
    cd $TOPDIR
done

rm -rf $tmpdir