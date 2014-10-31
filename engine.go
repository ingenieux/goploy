package goploy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

const (
	SERVICE          = "devtools"
	TERMINATOR       = "aws4_request"
	DATE_FORMAT      = "20060102"
	DATE_TIME_FORMAT = "20060102T150405"
	SCHEME           = "AWS4"
	AWS_ALGORITHM    = "HMAC-SHA256"
)

var (
	COMMITID_RE      = regexp.MustCompile("^[[:xdigit:]]{40}$")
	EINVALIDCOMMITID = errors.New("Invalid commitId")
)

type Deployment struct {
	accessKey       string
	secretKey       string
	applicationName string
	environmentName string
	region          string
	commitId        string
	time            time.Time
	date            string
	dateTime        string
}

func (d *Deployment) Region(newRegion string) {
	d.region = newRegion
}

func (d *Deployment) ApplicationName(newApplicationName string) {
	d.applicationName = newApplicationName
}

func (d *Deployment) EnvironmentName(newEnvironmentName string) {
	d.environmentName = newEnvironmentName
}

func (d *Deployment) CommitId(newCommitId string) error {
	if !COMMITID_RE.MatchString(newCommitId) {
		return EINVALIDCOMMITID
	}

	d.commitId = newCommitId
	return nil
}

func (d *Deployment) deriveKey() []byte {
	kDate := hash([]byte(SCHEME+d.secretKey), d.date)
	kRegion := hash(kDate, d.region)
	kService := hash(kRegion, SERVICE)
	kResult := hash(kService, TERMINATOR)

	return kResult
}

func (d *Deployment) GetPushURL() string {
	now := d.time.In(time.UTC)

	d.dateTime = now.Format(DATE_TIME_FORMAT)
	d.date = now.Format(DATE_FORMAT)

	host := fmt.Sprintf("git.elasticbeanstalk.%s.amazonaws.com", d.region)

	path := fmt.Sprintf("/v1/repos/%s/commitid/%s",
		hex.EncodeToString([]byte(d.applicationName)), hex.EncodeToString([]byte(d.commitId)))

	if "" != d.environmentName {
		path += fmt.Sprintf("/environment/%s", hex.EncodeToString([]byte(d.environmentName)))
	}

	scope := fmt.Sprintf("%s/%s/%s/%s", d.date, d.region, SERVICE, TERMINATOR)

	strToSign := fmt.Sprintf("%s-%s\n%s\n%s\n", SCHEME,
		AWS_ALGORITHM, d.dateTime, scope)

	strToSign += sha256hex(fmt.Sprintf("GIT\n%s\n\nhost:%s\n\nhost\n", path, host))

	key := d.deriveKey()

	digest := hash(key, strToSign)

	signature := hex.EncodeToString(digest)

	password := d.dateTime + "Z" + signature

	result := fmt.Sprintf("https://%s:%s@%s%s", d.accessKey, password, host, path)

	return result
}

func hash(kSecret []byte, obj string) []byte {
	mac := hmac.New(sha256.New, kSecret)

	mac.Write([]byte(obj))

	return mac.Sum(nil)
}

func sha256hex(s string) string {
	hash := sha256.New()

	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil))
}

func NewDeployment() (*Deployment, error) {
	accessKeyId, err := tryOrPanic("AWS_ACCESS_KEY_ID")

	if nil != err {
		return nil, err
	}

	secretAccessKey, err := tryOrPanic("AWS_SECRET_ACCESS_KEY")

	if nil != err {
		return nil, err
	}

	region := tryDefault("AWS_DEFAULT_REGION", "us-east-1")

	result := Deployment{
		accessKey: accessKeyId,
		secretKey: secretAccessKey,
		region:    region,
		time:      time.Now(),
	}

	return &result, nil
}

func tryOrPanic(env string) (string, error) {
	result := os.Getenv(env)

	if "" == result {
		return "", errors.New("Missing variable " + env)
	}

	return result, nil
}

func tryDefault(env, defaultValue string) string {
	result := os.Getenv(env)

	if "" == result {
		result = defaultValue
	}

	return result
}
