# github_webhook_action

github webhook 을 받아 필요한 액션(api 호출)을 수행합니다.

- build

```bash
go get ./...

# run for local test
go run .

# build
go build

# 실행
./main

```

## google app engine 사용 (비용 발생으로 2025.04 삭제)

```bash
# gcloud 설치 - mac
brew install google-cloud-sdk
# 또는 수동 설치
cd ~/workspace
# x86_64
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-477.0.0-darwin-x86_64.tar.gz
# arm64
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-477.0.0-darwin-arm.tar.gz
rm -rf google-cloud-sdk
tar zxvf google-cloud-cli-477.0.0-darwin-arm.tar.gz
# 컴포넌트 설치, golang 패키지 설치
gcloud components install app-engine-go

# gcloud 설치 - ubuntu
export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)"
echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update && sudo apt-get install google-cloud-sdk
sudo apt-get install google-cloud-sdk-app-engine-python google-cloud-sdk-app-engine-go google-cloud-sdk-datastore-emulator

# log
https://cloud.google.com/appengine/docs/standard/go/logs/
# app.yaml
https://cloud.google.com/appengine/docs/standard/go/config/appref

# google cloud 올리기전에 로컬에서 테스트 해볼 수 있다.
# 아래 명령을 실행해두면 .go 소스 수정때마다 자동 빌드 된다.
# google-cloud-sdk.tar.gz 로 설치했을 경우
~/workspace/google-cloud-sdk/bin/dev_appserver.py app.yaml --port 9999

# gcloud 인증(브라우저 열리고 로그인)
gcloud auth login

# google cloud 초기화
# url 링크 후 verification code 확인하여 입력
# 기존 프로젝트 또는 신규 프로젝트 생성 선택
# Compute Region and Zone 선택
gcloud init

# 다른 프로젝트 설정도 있다면 확인하고 default 변경
gcloud projects list
gcloud config set project github-webhook-action

# glcoud 구글 app engine 에 배포하기
# 배포 종료시 접속 가능한 url 이 표시된다.
# 배포전 아래 내용이 출력된다. 이상이 있다면 gcloud init 로 다시 설정하자.
# descriptor:      [/Users/ysoftman/workspace/github-webhook-action/app.yaml]
# source:          [/Users/ysoftman/workspace/github-webhook-action]
# target project:  [github-webhook-action]
# target service:  [default]
# target version:  [20190416t141405]
# target url:      [https://github-webhook-action.appspot.com]
# --verion 버전 명시
# --promote 현재 배포한 버전이 모든 트랙픽(100%)을 받도록 한다. 기존 버전의 인스턴스는 트랙픽 0% 이 된다.
# --stop-previous-version 새버전이 올라가면 기존 버전은 stop 하도록 한다.
# 현재시간으로 버전 설정
sed -i "" "s/^BuildTime=\".*/BuildTime=\"$(date '+%Y-%m-%d_%H:%M:%S_%Z')\"/" config.toml
gcloud app deploy ./app.yaml --version $(date '+%Y%m%d') --promote --stop-previous-version

# 빌드 로그
https://console.cloud.google.com/cloud-build

# 배포가 완료되면 확인
https://github-webhook-action.appspot.com

# 배포 후 접속 URL 확인 하기
gcloud app browse

# 앱 로그 확인
https://console.cloud.google.com/logs/viewer?project=github-webhook-action
gcloud app logs tail -s default

# instaces list
gcloud app instances list

# ssh 접속
gcloud app instances ssh default --server=default --version=20240604
```

## webhook 등록

- <https://github.com/ysoftman/test_code/settings/hooks> 에 등록
  - payloadURL: <https://github-webhook-action.appspot.com/v1/webhook>
  - secret: ysoftman
  - trgger(indiviual event): commit comment, pushes, pull requests, pull request reviews, pull request review comments ...등 트리거
