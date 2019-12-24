## About the app 
#
<h1 align="center"> <img width="300px" src="assets/symilogo.png" /></h1> 

[![CircleCI](https://circleci.com/gh/project-symi/backend-symi/tree/master.svg?style=shield)](https://circleci.com/gh/project-symi/backend-symi/tree/master) 

<img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />

Symi is a gamified platform for promoting positive feedback and company morale. Employees get points for sending anonymous feedback directly to their CEO. CEO's can have a top-level overview of company sentiment in the form of data visualization personalized to company changes to employees.


## Additional Info
This is the backend server implementation for SYMI.

To test out the working version or learn more about the front end, please check out the demo link above or see the frontend [Readme](https://github.com/project-symi/frontend-symi). 
### ✨ [Demo](https://www.symi.dev)

## Technology

<img alt="Technology logo list" src="./assets/technologies.png">

## Running the backend server

### *Docker*
<img alt="Docker logo" src="https://res.cloudinary.com/teepublic/image/private/s--bLlrDo5F--/c_crop,x_10,y_10/c_fit,w_1109/c_crop,g_north_west,h_1260,w_1260,x_-76,y_-165/co_rgb:ffffff,e_colorize,u_Misc:One%20Pixel%20Gray/c_scale,g_north_west,h_1260,w_1260/fl_layer_apply,g_north_west,x_-76,y_-165/bo_157px_solid_white/e_overlay,fl_layer_apply,h_1260,l_Misc:Art%20Print%20Bumpmap,w_1260/e_shadow,x_6,y_6/c_limit,h_1134,w_1134/c_lpad,g_center,h_1260,w_1260/b_rgb:eeeeee/c_limit,f_jpg,h_630,q_90,w_630/v1521449955/production/designs/2490921_0.jpg" width="100" />

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
<img alt="Golang logo" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/512px-Go_Logo_Blue.svg.png" width="100px">

Directly running the go application is also possible with  
```sh
go run ./application.go
```

## Deploying

Symi-backend contains 2 setups that can be utilized - AWS and Heroku
In both cases, please make sure to add the following Environment variables:

`DB_INFO: the full link to the Database you plan to use for the app <user>:<password>@tcp(<link_to_db>)/<db_name>`<br>
`SIGNING_KEY: <string> - will be used as SALT to hash and check your passwords for storage.`

Make sure to run migrations (and seed some data) from `/app/migrations` and `/app/seeder`.

### *Heroku*
<img alt="Heroku logo" src="https://redislabs.com/wp-content/uploads/2016/11/logo-square-heroku.png" width="100px">

To deploy the backend server Heroku, create an app via the CLI or on [Heroku](https://heroku.com).
Fork your version of the backend-SYMI repository.
Link your created app with the forked github repository.
Click "Deploy" in settings!

Based on Heroku.yml, a docker image will be created. Heroku will then start up a container in your app.

### *AWS*
<img alt="AWS logo" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/1024px-Amazon_Web_Services_Logo.svg.png" width="100px">

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

### AWS Workflow

<img alt="AWS CloudFormation" src="./assets/aws/AWS_CloudFormation.png" width="100px">
<img alt="AWS CodePipeline logo" src="./assets/aws/AWS_CodePipeline.png" width="100px">
<img alt="AWS CodePipeline logo" src="./assets/aws/AWS_CodeBuild.png" width="100px">
<img alt="AWS ECS Fargate logo" src="./assets/aws/AWS_ECS_Fargate.png" width="100px">
<img alt="AWS RDS logo" src="./assets/aws/AWS_RDS.png" width="100px">

The main steps of the deployment process are as follows:

1. User manually deploys the initial Cloudformation template containing the CodePipeline infrastructure (`CodePipeline.yml`).
2. Cloudformation starts the deployment of the AWS CodePipeline infrastructure.
3. After the Cloudformation deployment, CodePipeline pulls the source code from Github in the source stage.
4. CodePipeline builds the docker image and pushes to the ECR in the build stage using AWS CodeBuild service.
5. CodePipeline starts the deployment of the Cloudformation template (`Fargate-Cluster.yml`) containing Fargate ECS Cluster in the deploy stage.

Simplified structure can be seen below:
<img alt="SYMI CI/CD Pipeline flow" src="./assets/aws/AWS_Schema.png">


For more details details and information, please check out the following article: [AWS Cloudformation Managed (Medium.com)](https://medium.com/swlh/aws-cloudformation-managed-complete-ecs-infrastructure-including-ci-cd-pipeline-from-github-to-ecs-b833bb44e01c)

## Infrastructure

For this project we implemented a clean architecture ([Onion Architecture](https://www.thinktocode.com/2018/08/21/onion-architecture-skeleton-example/)).

## Database Structure

To match the complex data flow of SYMI information, the below structure was created.

<img alt="SYMI database structure" src="./assets/SYMI_db_structure.png" />

Furthermore, due to the potentially large amount of user actions, instead of using [ORM](https://en.wikipedia.org/wiki/Object-relational_mapping) for DB interactions, the read/write operations are done using pure SQL, to achieve faster handling speed.

## Future Features

* Utilize Go routines and/or channels to further speed up data handling.

* Breakdown the monolith code into microservices (employee actions would be loaded more than CEO, so scaling only that part would be easier).

* Create functionality to support the future features mentioned on the frontend Readme:
  * Custom Assignments
  * Employee Slack Notifications
  * Ivite Calendar Integration
  * Points Animations
  * Badges & Rewards


## Contributors
<table>
 <tr>
    <td align="center">
      <a href="https://github.com/miniengineer">
        <img src="./assets/headshots/mini.png"" width="200px;"/><br />
        <sub><b>Mini</b></sub>
      </a><br />Tech Lead
    </td>
    <td align="center"><a href="https://github.com/FuyuByakko"><img src="./assets/headshots/igor.png" width="200px;"/><br /><sub><b>Igor</b></sub></a><br />Fullstack</td>
    <td align="center"><a href="https://github.com/steffieharner"><img src="./assets/headshots/steffie.png" width="200px;"/><br /><sub><b>Steffie Harner</b></sub></a><br />Design/Frontend</td> 
    <td align="center">
      <a href="https://github.com/Yukio0315"><img src="./assets/headshots/yukio.png" width="200px;"/><br />
        <sub>
          <b>Yukio Ueda</b>
        </sub>
      </a><br />
      Backend
    </td>
 </tr>
</table>

## Show your support

Give a ⭐️ if you like our stuff!

## 📝 License

This project is [ISC](https://github.com/project-symi/frontend-symi/license.md) licensed.
