# Jackal
[![Build Status](https://travis-ci.com/astoliarov/jackal.svg?branch=master)](https://travis-ci.com/astoliarov/jackal)

Image resizing proxy. Simple opensourced replacement for [CloudFront + AWS image resizing solution](https://aws.amazon.com/blogs/networking-and-content-delivery/resizing-images-with-amazon-cloudfront-lambdaedge-aws-cdn-blog/).

### Features
- On-the-fly image processing: Jackal fetch target image and resize it on demand
- Caching (not implemented right now. See TODO section): Jackal acts like cache for processed images.
- Two types of crop: exact and by ratio

### How to use it
Build a Docker image:
```
docker build . -t jackal
```
Start a container:
```
docker run --name awesome-jackal -p 0.0.0.0:3000:3000 jackal
```
Now you can perform requests:
```
wget 'http://localhost:3000/api/v1/crop?height=300&width=400&url=https%3A%2F%2Fsource.unsplash.com%2Frandom%2F800x600'
```

Api description placed in [api/openapi.yaml](https://github.com/astoliarov/jackal/blob/master/api/openapi.yaml)

### TODO:
 - [ ] Cache for processed images (store result in different storages: machine, S3)
 - [ ] More option to Crop endpoint (different crop positions)
 - [ ] Resize endpoint

### Contributing

Please fork the project, clone your repository and add the original repo as an upstream remote to keep yours in sync.
For small fixes (typo and such), you can work on master.


