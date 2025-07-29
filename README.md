# CI Pipeline with Jenkins for Go Web App

## 👨‍💻 Author

**Chennuri Akhil Varma**  
Kluisz.ai

---

## 📌 Project Overview

This repository contains a basic Continuous Integration (CI) setup using **Jenkins** for a dummy **Go web application**. It includes:

- Source code for the Go application
- Jenkins pipeline configuration (`Jenkinsfile`)
- Kubernetes deployment and service manifests
- A Dockerfile to containerize the app

The purpose is to demonstrate end-to-end CI using Jenkins, containerization using Docker, and Kubernetes deployment — all in a local environment.

---

## 💻 UI Details

The Go application is a simple web service with a basic UI built using plain HTML. The app serves a welcome message and can be extended as needed.

- Route: `/`
- Method: `GET`
- Output: `Welcome to CI/CD Jenkins Go App`

---

## 🔁 CI Process Flow

Here’s how the Jenkins CI pipeline works:

1. **Clone Stage**

   - Jenkins pulls the code from GitHub using webhooks or polling.

2. **Build Stage**

   - The Go application is compiled.

3. **Test Stage**

   - Unit tests (if added) are executed using `go test`.

4. **Docker Build Stage**

   - Jenkins builds a Docker image from the `Dockerfile`.

5. **Docker Push Stage**

   - Image is pushed to Docker Hub (optional for local testing).

6. **Deploy Stage**
   - Kubernetes deployment and service manifests are applied using `kubectl`.

> This pipeline runs inside Jenkins and can be customized for additional steps (e.g., linting, security scans, etc.)

---

## 🛠 How to Set Up Locally

> Make sure you have these installed:
>
> - Git
> - Docker
> - Minikube or Kubernetes
> - Jenkins (run inside Docker)

---

### 🚀 Clone the Repository

```bash
git clone https://github.com/ChennuriAkhilvarma/go-app.git
cd go-app
🧱 Run Jenkins in Docker

docker run -d \
  -u root \
  --name jenkins \
  -p 8080:8080 -p 50000:50000 \
  -v jenkins_home:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkins/jenkins:lts


🔑 Get Jenkins Admin Password
docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword


Open Jenkins UI: http://localhost:8080

Paste the password to complete setup.



🧩 Install Required Plugins

Inside Jenkins → Manage Jenkins → Plugins:

Docker
Kubernetes CLI
Git
Pipeline


📝 Add Jenkinsfile

Ensure your Jenkinsfile is at the root of your repo. It should define all stages like:

pipeline {
  agent any
  stages {
    stage('Clone') {
      steps {
        git 'https://github.com/ChennuriAkhilvarma/ci-pipeline-go-app.git'
      }
    }
    ...
  }
}



☸️ Kubernetes Deployment (via Jenkins)

Your repo should contain:

deployment.yaml
service.yaml

They will be applied during the Deploy stage via kubectl apply.

✅ Final Output
Once the pipeline completes, your app will be running inside a Kubernetes pod.
You can access the service using:

minikube service go-app-service

🧾 Folder Structure

.
├── Dockerfile
├── Jenkinsfile
├── deployment.yaml
├── service.yaml
├── main.go
└── README.md

```
