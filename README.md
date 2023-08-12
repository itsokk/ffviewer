# ffviewer

ffviewer is a simple and easy-to-use cross-platform farbfeld viewer made with raylib and Go.

# What is farbfeld?

[Farbfeld](https://tools.suckless.org/farbfeld/) is a lossless image format which is easy to parse, pipe and compress. It has the following format:
| Bytes  | Description                                             |
|--------|---------------------------------------------------------|
| 8      | "farbfeld" magic value                                  |
| 4      | 32-Bit BE unsigned integer (width)                      |
| 4      | 32-Bit BE unsigned integer (height)                     |
| [2222] | 4x16-Bit BE unsigned integers [RGBA] / pixel, row-major |

# Requirements

[Go](https://go.dev/)

The raylib-go requirements for each OS are listed [here](https://github.com/gen2brain/raylib-go#requirements)