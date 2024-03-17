# DistributedRaytracer

The objective of this project is to create a distributed raytracer service capable of rendering animations. The service will achieve this by splitting the frames of the animation into smaller tasks and distributing them to multiple workers. These workers will render the frames and send them back to the service, which will then combine them to create the final animation.

## Design

![Design](docs/design.svg)

## Example

![Example](docs/render.png)

## Build

```bash
go build -o raytracer
./raytracer
```

## Usage

```bash
Usage of ./raytracer:
  -cpu int
        number of CPU to use (default to number of CPU available) (default 14)
  -h int
        height in pixel (default 400)
  -i string
        path to file for world definition in json format (default "assets/world.json")
  -o string
        path to file for saving (do not save if not defined) (default "output.png")
  -r int
        number of rays per pixel (default 100)
  -seed int
        seed for random number generator (default 2024)
  -w int
        width in pixel (default 800)
```

## Reference

- [Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)
