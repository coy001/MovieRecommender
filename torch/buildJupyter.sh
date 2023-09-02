docker build --no-cache \
             -f create-user.dockerfile \
             --build-arg MY_UID="$(id -u)" \
             --build-arg MY_GID="$(id -g)" \
             --build-arg USER=demo \
             --build-arg HOME=/home/demo \
             -t tverous/pytorch-notebook:user \
             .
