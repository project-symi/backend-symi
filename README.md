## About the app 
#
<h1 align="center"> <img width="300px" src="assets/symilogo.png" /></h1>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />

Creating transparency between CEO and bottom line.
For more information about testing out the demo, please check the front-end readme.

### 🏠 [Homepage](https://github.com/project-symi/frontend-symi#readme)

### ✨ [Demo](<add demo url>)

## Running the backend server

### *Docker*
SYMI backend makes use of Docker to create and image of the app and be able to deploy a docker container quickly and with easy

```sh
docker build -t symi-backend
//symi-backend can be replaced with your desired image name
```

```sh
docker run --rm -it -p 5000:8080 -e PORT=8080 symiback-end
```
`-p 5000` -> determines which port your container will be accessible on<br>
`:8080` -> which port the app is listening on (inside the container).<br>
By default the app listening port is 8080, however it can be changed by passing in a PORT environment variable (-e PORT=####).

### *Makefile*
If needed, it is possible to utilize Makefile, instead of Docker.
If you have Make installed, you can run the app with  
```sh
make run
```

### *Go*
Directly running the go application is also possible with  
```sh
go run ./application.go
```

## Deploying

Symi-backend contains 2 setups that can be utilized - AWS and Heroku
In both cases, please make sure to add the following Environment variables:

`DB_INFO: the full link to the Database you plan to use for the app <user>:<password>@tcp(<link_to_db>)/<db_name>`

`SIGNING_KEY: <string> - will be used as SALT to hash and check your passwords for storage.`

Make sure to run migrations (and seed some data) from `/app/migrations` and `/app/seeder`

### *Heroku*

To deploy the backend server Heroku, create an app via the CLI or on [Heroku](https://heroku.com).
Fork your version of the backend-SYMI repository.
Link your created app with the forked github repository.
Click "Deploy" in settings!

Based on Heroku.yml, a docker image will be created. heroku will then start up a container in your app.

### *AWS*

AWS Deployment will require some prior setup.
The first step will be to create the initial CI/CD Pipeline that will be building the Docker image and deploying the rest of the infrastructure.
This step has to be performed manually.

Either via the AWS CLI tool, or using the [AWS Cloudformation website](https://aws.amazon.com/cloudformation/), create a new stack using the ready template.
Template: `/cloudformation/codepipeline.yml`

* Please make sure to add a Personal Access Token to your account allowing AWS CFN access to the repository.
* Change the current template info GithubUserName/GithubRepo/GithubBranch to match your fork.

Once the template is created, With every push to the branch that was specified in GithubBranch, the AWS Codepipeline will build the Docker image, store it into an S3 bucket, and finally trigger the infrastructure build using ECS Fargate.

Once again, by default Fargate will create 2 tasks running in 2 different Subnets and create a loadbalancer that connects them both.
To get access to the server, plese find the DNS adress in the ECS Instances -> ElasticLoadBalancers.

## Running the backend server

## Database Structure

To match the complex data flow of SYMI information, the below structure was created.

<img alt="SYMI database structure" src="./assets/SYMI_db_structure.png" />


## Contributors

<table height="500px">
 <tr>
    <td align="center"><a href="https://github.com/miniengineer"><img src="./assets/headshots/mini.png"" width="200px;"/><br /><sub><b>Mini</b></sub></a><br />Tech Lead</td>
    <td align="center"><a href="https://github.com/FuyuByakko"><img src="./assets/headshots/igor.png" width="200px;"/><br /><sub><b>Igor</b></sub></a><br />Fullstack</td>
    <td align="center"><a href="https://github.com/steffieharner"><img src="./assets/headshots/steffie.png" width="200px;"/><br /><sub><b>Steffie Harner</b></sub></a><br />Design/Frontend <p>
  <a href="https://twitter.com/steffieharner" target="_blank">
    <img alt="Twitter: steffieharner" src="https://img.shields.io/twitter/follow/steffieharner.svg?style=social" />
  </a>
</p></td> 
    <td align="center"><a href="https://github.com/Yukio0315"><img src="./assets/headshots/yukio.png" width="200px;" alt=""/><br /><sub><b>Yukio Ueda</b></sub></a><br /> Backend <p>
  <a href="https://twitter.com/SnowSnowManMan" target="_blank">
    <img alt="Twitter: SnowSnowManMan" src="https://img.shields.io/twitter/follow/SnowSnowManMan.svg?style=social" />
  </a>
</p>
  </tr>
</table>

## Show your support

Give a ⭐️ if you like our stuff!

## 📝 License

This project is [ISC](https://github.com/project-symi/frontend-symi/license.md) licensed.
