docker run --rm -it  \
           -p 8888:8888  \
	   -v '/Users/luanqinglin/desktop/MovieRecommender/torch':'/Users/luanqinglin/desktop/MovieRecommender/torch' -w '/Users/luanqinglin/desktop/MovieRecommender/torch' \
           -e JUPYTER_TOKEN=passwd  \
           tverous/pytorch-notebook:latest
