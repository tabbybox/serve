#!/bin/bash
for os in darwin linux windows
do
    for arch in 386 amd64 arm64
    do  
        if [[ $os == "darwin" &&  $arch == "386" ]]; then
            echo Skipping Apple-x86
            continue
        fi
        echo Building for $os-$arch
        GOOS=$os GOARCH=$arch go build -o build/serve-$os-$arch &
        # echo $os-$arch
    done
done
wait
echo "Compile complete"