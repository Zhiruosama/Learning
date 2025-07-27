// pngquant 是一个用于压缩 PNG 图像文件的工具。它可以显著减小 PNG 文件的大小
// 同时保持图像质量和透明度。通过减小文件大小，可以提高网页加载速度，并节省存储空间

import { exec } from 'child_process'

exec('pngquant 73kb.png --output test.png')
//quality范围为0-100
exec('pngquant 73kb.png --quality=100 --output test.png')
//speed范围0-10 数字越小越慢 质量越高
exec('pngquant 73kb.png --speed=10 --quality=100 --output test.png')