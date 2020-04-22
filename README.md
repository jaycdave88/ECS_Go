### Instructions to setup APM for Golang within AWS ECS/Fargate: 

Setup Locally: 
- Run the following to start the Golang application, (main.go): 
   1. `docker build -t <container_tag> .`
   2. `docker run -p 9090:9090 -it <container_id>`
   3. Navigate to `localhost:9090/<test name>`

Setup in AWS ECS: 

   1. Create a Datadog agent daemon service. Instructions can be found here: https://docs.datadoghq.com/integrations/amazon_ecs/?tab=python#web-ui
         1.a. (Note. The `Task_Definition_Datadog_agent.json` can be used. Please update to include your Datadog API key.)
   2. Create/modify GoApplication's ECS task definition with `Task_definition_GoApplication.json`. The important piece is to include the application **portMappings** for **9090**.
