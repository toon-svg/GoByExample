from golang:1.16-buster
RUN apt update 
RUN apt install -y vim \
git \
curl
RUN git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime
RUN git clone https://github.com/preservim/nerdtree.git ~/.vim/pack/vendor/start/nerdtree
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
COPY .vimrc /root/.vimrc