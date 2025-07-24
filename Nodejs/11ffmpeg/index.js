// FFmpeg 的主要功能和特性：

// 格式转换：FFmpeg 可以将一个媒体文件从一种格式转换为另一种格式，支持几乎所有常见的音频和视频格式，包括 MP4、AVI、MKV、MOV、FLV、MP3、AAC 等。
// 视频处理：FFmpeg 可以进行视频编码、解码、裁剪、旋转、缩放、调整帧率、添加水印等操作。你可以使用它来调整视频的分辨率、剪辑和拼接视频片段，以及对视频进行各种效果处理。
// 音频处理：FFmpeg 可以进行音频编码、解码、剪辑、混音、音量调节等操作。你可以用它来提取音频轨道、剪辑和拼接音频片段，以及对音频进行降噪、均衡器等处理。
// 流媒体传输：FFmpeg 支持将音视频流实时传输到网络上，可以用于实时流媒体服务、直播和视频会议等应用场景。
// 视频处理效率高：FFmpeg 是一个高效的工具，针对处理大型视频文件和高分辨率视频进行了优化，可以在保持良好质量的同时提供较快的处理速度。
// 跨平台支持：FFmpeg 可以在多个操作系统上运行，包括 Windows、MacOS、Linux 等，同时支持多种硬件加速技术，如 NVIDIA CUDA 和 Intel Quick Sync Video。

const { execSync } = require('child_process')
//1.基本格式转化 avi mp4 gif
// execSync('ffmpeg -i test.mov test.avi',{stdio:'inherit'})
//{stdio:'inherit'}参数可以让其输出转换过程中的详细信息

//2.音频提取
// execSync('ffmpeg -i test.mov test.mp3',{stdio:'inherit'})

//3.裁剪视频
// execSync('ffmpeg -ss 10 -to 20 -i test.mov test.mp4',{stdio:'inherit'})

//4.加水印 -vf
// execSync('ffmpeg -i test.mov -vf drawtext=text="CSCS":fontsize=30:x=10:y=10:fontcolor=white test2.mp4', { stdio: 'inherit' })

//5.删除水印
// execSync('ffmpeg -i test2.mp4 -vf delogo=w=120:h=30:x=10:y=10 test3.mp4', { stdio: 'inherit' })

