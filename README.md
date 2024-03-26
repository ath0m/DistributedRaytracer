# DistributedRaytracer

The objective of this project is to create a distributed raytracer service capable of rendering animations. The service will achieve this by splitting the frames of the animation into smaller tasks and distributing them to multiple workers. These workers will render the frames and send them back to the service, which will then combine them to create the final animation.

## Design

![Design](docs/design.svg)

## Example

![Example](docs/render.png)

## Usage

To generate a new render and save it to `output.png` file on local machine, start application and perform a POST request:

```bash
docker compose up -d
curl -X POST http://localhost:8090/render -v -d '{"width":800, "height": 400, "raysperpixel": 10, "seed": 2024}' --output output.png
curl -X POST http://localhost:8090/render -v -d '{"width":800, "height": 400, "raysperpixel": 10, "seed": 2024, "world": {"camera":{"origin":{"X":13,"Y":2,"Z":3},"lowerLeftCorner":{"X":2.8254931764402573,"Y":-1.2262841980681716,"Z":4.2712604900308655},"horizontal":{"X":1.5859519159914772,"Y":0,"Z":-6.872458302629735},"vertical":{"X":-0.5094205020606202,"Y":3.4875711294919385,"Z":-0.11755857739860466},"u":{"X":0.22485950669875845,"Y":0,"Z":-0.97439119569462},"v":{"X":-0.14445336159384606,"Y":0.9889499370655616,"Z":-0.0333353911370414},"lensRadius":0.05},"objects":[{"center":{"X":0,"Y":-1000,"Z":0},"radius":1000,"material":{"type":"Lambertian","albedo":{"R":0.5,"G":0.5,"B":0.5}}},{"center":{"X":0,"Y":1,"Z":0},"radius":1,"material":{"type":"Dielectric","refIdx":1.5}},{"center":{"X":-4,"Y":1,"Z":0},"radius":1,"material":{"type":"Lambertian","albedo":{"R":0.4,"G":0.2,"B":0.1}}},{"center":{"X":4,"Y":1,"Z":0},"radius":1,"material":{"type":"Metal","albedo":{"R":0.7,"G":0.6,"B":0.5},"fuzz":0}}]}}' --output output.png
```

## Reference

- [Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)
