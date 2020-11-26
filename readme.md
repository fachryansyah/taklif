# TAKLIF
Based on Indonesia Dictionary Taklif means "submission of heavy burdens (work, duties, etc.) (to someone)" the point is heavy burdens when Deploy your application. with Taklif your deployments can be simple like serverless but running on your VPS / Dedicated Server 🔥

# Goals 🚀
Make your deployment simple like a serverless, just push your code and integrate your GIT then Taklif will running your app.

# Guide 📜
## Folder Structure
### Client
Client folder is place for GUI (Graphical User Interface) to manage, config of apllication
### Public
A place for storing production build of Client
### Handlers
Collection of Rest API controller for Client
### PM
PM is Proccess Manager that can manage your application, create app, check status, kill, delete and etc.
### Routes
Route URL of Rest API from controller

# Todo 🗒
1. Web Client
- [ ] Setup VueJs App & Tailwind in Client folder
- [ ] Slicing Client page
2. Web Services
- [ ] Setup Backend REST API with Go Fiber
- [ ] Create Mock REST API
- [ ] Create app management
- [ ] Create custom web hook for Github & Gitlab
3. Proccess Manager
- [ ] Setup pm
- [ ] Make pm can (running, kill, delete, restart, logging) go application
- [ ] Make pm can (running, kill, delete, restart, logging) nodejs application with nvm
- [ ] Make pm can (running, kill, delete, restart, logging) php application with nginx
- [ ] Create API for monitoring the application

# Contributor 👑
