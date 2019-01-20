VERSION="$1"
RELEASE="release/$VERSION"

make build

docker build -f manifests/Dockerfile -t us.gcr.io/bitja-193417/simple-proxy:$VERSION .
docker push us.gcr.io/bitja-193417/simple-proxy:$VERSION

git checkout -b $RELEASE
git tag -a $VERSION -m "release ${VERSION}"
git push origin $VERSION
git checkout master
git merge $RELEASE
git push origin master
