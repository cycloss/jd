#! /bin/sh

if [[ -z "$JD_INSTALL_PATH" ]]; then
    JD_INSTALL_PATH='/usr/local/bin'
fi

echo "Installing jd binary to $JD_INSTALL_PATH"

go build -o "$JD_INSTALL_PATH/jd" ./...

if [[ $? -ne 0 ]]; then
    echo "Failed to install jd"
    exit 1
fi

JD_CONF_PATH='/usr/local/etc/jd'

echo "Installing jd config to $JD_CONF_PATH"

cp -R etc/jd/. $JD_CONF_PATH

if [[ $? -ne 0 ]]; then
    echo "Failed to install jd"
    exit 1
fi

echo  "Successfully installed jd"