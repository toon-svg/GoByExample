autocmd VimEnter * NERDTree
call plug#begin()
Plug 'preservim/nerdtree'
Plug 'ryanoasis/vim-devicons'
Plug 'sainnhe/gruvbox-material'
call plug#end()
set background=dark
colorscheme gruvbox-material
nnoremap <leader>n :NERDTreeFocus<CR>
nnoremap <C-n> :NERDTreeToggle<CR>
nnoremap <C-f> :NERDTreeFind<CR>
set nu