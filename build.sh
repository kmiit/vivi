set -e

buildDate="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')

version=0.0.1

ldflags="\
-w -s \
-X 'github.com/kmiit/vivi/utils/version.BuildDate=$buildDate' \
-X 'github.com/kmiit/vivi/utils/version.GoVersion=$goVersion' \
-X 'github.com/kmiit/vivi/utils/version.Version=$version' \
"

SHELL_DIR=$(dirname $(readlink -f "$0"))
OUT_DIR=$SHELL_DIR/out
SOURCE_DIR=$SHELL_DIR/server
cd $SHELL_DIR

if [[ ! -d $OUT_DIR ]]; then
    mkdir $OUT_DIR
fi

BuildAndroidArm64(){
    fileName=vivi-$version-android-arm64
    cd $SOURCE_DIR
    GOOS=android GOARCH=arm64 go build -ldflags="$ldflags" -o $OUT_DIR/$fileName
}

BuildLinuxAmd64() {
    fileName=vivi-$version-linux-amd64
    cd $SOURCE_DIR
    GOOS=linux GOARCH=amd64 go build -ldflags="$ldflags" -o $OUT_DIR/$fileName
}

BuildLinuxArm64() {
    fileName=vivi-$version-linux-arm64
    cd $SOURCE_DIR
    GOOS=linux GOARCH=arm64 go build -ldflags="$ldflags" -o $OUT_DIR/$fileName
}

BuildLocalMachine() {
    fileName=vivi-$version
    cd $SOURCE_DIR
    go build -ldflags="$ldflags" -o $OUT_DIR/$fileName
}

if [[ "$1" == android && "$2" == arm64 ]]; then
BuildAndroidArm64
else if [[ "$1" == linux && "$2" == amd64 ]]; then
BuildLinuxAmd64
else if [[ "$1" == linux && "$2" == arm64 ]]; then
BuildLinuxArm64
else
BuildLocalMachine
fi;fi;fi
