if exists('g:loaded_renata')
  finish
endif
let g:loaded_renata = 1

function! s:RequireRenata(host) abort
  return jobstart(['renata.nvim'], { 'rpc': v:true })
endfunction

call remote#host#Register('renata.nvim', '0', function('s:RequireRenata'))
call remote#host#RegisterPlugin('renata.nvim', '0', [
  \ {'type': 'command', 'name': 'Http', 'sync': 1, 'opts': {'nargs': '?'}},
  \ ])
