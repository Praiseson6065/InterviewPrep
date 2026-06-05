# Part 4: MLOps, Cloud, Docker, Kubernetes & System Design

---

## ⚙️ MLOps

### Beginner

**Q1: What is MLOps?**
- Practices for deploying and maintaining ML models in production **reliably and efficiently**
- Combines ML + DevOps + Data Engineering
- Key pillars: versioning, CI/CD, monitoring, reproducibility, automation

**Q2: What is the ML lifecycle?**
```
Data Collection → Data Prep → Feature Engineering → Model Training
→ Evaluation → Deployment → Monitoring → Retraining (loop)
```

**Q3: What is experiment tracking?**
- Record hyperparameters, metrics, artifacts for every training run
- Tools: **MLflow**, Weights & Biases (W&B), Neptune, TensorBoard
```python
import mlflow
with mlflow.start_run():
    mlflow.log_param("learning_rate", 0.001)
    mlflow.log_param("epochs", 50)
    mlflow.log_metric("accuracy", 0.95)
    mlflow.sklearn.log_model(model, "model")
```

### Intermediate

**Q4: Explain model serving patterns.**

| Pattern | Description | Use Case |
|---------|-------------|----------|
| Batch | Run predictions on schedule | Reports, recommendations |
| Real-time (REST) | Synchronous API calls | User-facing predictions |
| Streaming | Process events in real-time | Fraud detection |
| Edge | Run model on device | Mobile, IoT |

**Q5: What is a feature store?**
- Centralized repository for storing and serving ML features
- Ensures **consistency** between training and serving
- Tools: **Feast**, Tecton, Hopsworks
- Features: point-in-time correctness, online/offline stores, feature versioning

**Q6: Explain CI/CD for ML.**
```
Code Change → Unit Tests → Data Validation → Model Training
→ Model Validation (performance thresholds) → Staging Deploy
→ A/B Test / Shadow Mode → Production Deploy → Monitor
```
- **Data validation**: schema checks, distribution drift (Great Expectations)
- **Model validation**: performance ≥ baseline, fairness checks, latency constraints

### Advanced

**Q7: What is model monitoring and drift detection?**
- **Data drift**: input distribution changes (e.g., new user demographics)
- **Concept drift**: relationship between input and output changes
- **Detection**: statistical tests (KS test, PSI), monitoring dashboards
- **Response**: retrain, rollback, alert
- Tools: Evidently AI, WhyLabs, Arize

**Q8: Explain model versioning and registry.**
```
Model Registry (e.g., MLflow):
  model_name: "fraud_detector"
    ├── Version 1 (Staging)   - accuracy: 0.92
    ├── Version 2 (Production) - accuracy: 0.95
    └── Version 3 (Archived)   - accuracy: 0.91
```
- Tracks lineage: which data, code, and params produced which model
- Enables rollback, A/B testing, canary deployments

**Q9: Explain A/B testing for ML models.**
- Split traffic between current model (control) and new model (treatment)
- Measure business metrics (not just ML metrics)
- Statistical significance testing before full rollout
- **Shadow mode**: run new model in parallel without serving its results (compare offline)

---

## ☁️ Cloud Platforms

### Beginner

**Q1: Compare major cloud AI/ML services.**

| Service | GCP | AWS | Azure |
|---------|-----|-----|-------|
| ML Platform | Vertex AI | SageMaker | Azure ML |
| LLM API | Gemini API | Bedrock | Azure OpenAI |
| Compute | GCE / GKE | EC2 / EKS | AKS / VMs |
| Storage | GCS | S3 | Blob Storage |
| Serverless | Cloud Functions | Lambda | Azure Functions |
| Data Warehouse | BigQuery | Redshift | Synapse |

**Q2: What is the difference between IaaS, PaaS, and SaaS?**
- **IaaS** (Infrastructure): VMs, networking → EC2, GCE
- **PaaS** (Platform): managed runtime → App Engine, Elastic Beanstalk
- **SaaS** (Software): ready-to-use → Gmail, Salesforce

**Q3: Explain object storage (S3/GCS/Blob).**
- Store unstructured data (files, images, models) as objects
- Flat namespace with bucket/key structure
- Highly durable (11 nines), scalable, cheap
- Access via REST APIs, SDKs, CLI

### Intermediate

**Q4: How would you design a cloud ML training pipeline?**
```
GCP Example:
  Cloud Storage (data) → Vertex AI Pipelines → 
  Training on GPU VMs → Model Registry → 
  Vertex AI Endpoints (serving) → Cloud Monitoring

AWS Example:
  S3 (data) → SageMaker Pipelines → 
  Training on ml.p3 instances → Model Registry →
  SageMaker Endpoints → CloudWatch
```

**Q5: What is IAM and why is it critical?**
- Identity and Access Management — controls WHO can do WHAT on WHICH resources
- **Principle of least privilege**: grant minimum required permissions
- Concepts: users, roles, policies, service accounts
- Critical for: API keys, model endpoints, data access, cost control

**Q6: Explain serverless computing for AI.**
- **Cloud Functions / Lambda**: event-driven, auto-scaling, pay-per-invocation
- Good for: lightweight inference, preprocessing, webhooks
- Limitations: cold starts, execution time limits, memory constraints
- **Cloud Run / Fargate**: container-based serverless (better for ML models)

### Advanced

**Q7: How do you optimize cloud costs for ML workloads?**
- **Spot/Preemptible instances**: 60-90% cheaper for training (handle interruptions)
- **Auto-scaling**: scale serving endpoints based on traffic
- **Right-sizing**: match GPU type to workload (T4 for inference, A100 for training)
- **Storage tiers**: hot → cold → archive for infrequent data
- **Reserved instances**: commit for 1-3 years for steady workloads

---

## 🐳 Docker

### Beginner

**Q1: What is Docker and why use it for ML?**
- Containerization platform — packages code + dependencies into portable units
- **Reproducibility**: same environment everywhere (dev, staging, prod)
- **Isolation**: no dependency conflicts between projects
- **Consistency**: "works on my machine" problem solved

**Q2: Explain Dockerfile basics.**
```dockerfile
FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
```

**Q3: What is the difference between image and container?**
- **Image**: read-only template (blueprint) — like a class
- **Container**: running instance of an image — like an object
- `docker build` → creates image
- `docker run` → creates container from image

### Intermediate

**Q4: How do you optimize Docker images for ML?**
```dockerfile
# Multi-stage build
FROM python:3.11 AS builder
COPY requirements.txt .
RUN pip install --user -r requirements.txt

FROM python:3.11-slim
COPY --from=builder /root/.local /root/.local
COPY . .
CMD ["python", "serve.py"]
```
- Use slim/alpine base images
- Multi-stage builds to reduce final size
- Layer caching: put infrequently changing layers first
- `.dockerignore` to exclude data files, notebooks, `.git`

**Q5: Explain Docker Compose.**
```yaml
version: '3.8'
services:
  api:
    build: ./api
    ports: ["8000:8000"]
    environment:
      - MODEL_PATH=/models/latest
    depends_on: [vectordb]
  
  vectordb:
    image: qdrant/qdrant
    ports: ["6333:6333"]
    volumes: ["qdrant_data:/qdrant/storage"]

volumes:
  qdrant_data:
```

**Q6: How do you handle GPU access in Docker?**
```bash
# NVIDIA Container Toolkit required
docker run --gpus all -it pytorch/pytorch:latest python -c "
import torch; print(torch.cuda.is_available())
"

# Specify GPU count
docker run --gpus '"device=0,1"' my-training-image
```

---

## ☸️ Kubernetes

### Beginner

**Q1: What is Kubernetes and why use it?**
- Container orchestration platform — automates deployment, scaling, management
- **Why for ML**: scale model serving, manage GPU resources, rolling updates
- Key concepts: Pods, Services, Deployments, ConfigMaps, Secrets

**Q2: Explain Pods, Deployments, and Services.**
```yaml
# Deployment - manages replicas
apiVersion: apps/v1
kind: Deployment
metadata:
  name: model-server
spec:
  replicas: 3
  selector:
    matchLabels:
      app: model-server
  template:
    spec:
      containers:
      - name: model
        image: my-model:v2
        resources:
          limits:
            nvidia.com/gpu: 1
            memory: "4Gi"
---
# Service - stable endpoint
apiVersion: v1
kind: Service
metadata:
  name: model-service
spec:
  selector:
    app: model-server
  ports:
  - port: 80
    targetPort: 8000
  type: LoadBalancer
```

### Intermediate

**Q3: Explain Horizontal Pod Autoscaler (HPA).**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: model-server
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

**Q4: How do you do rolling updates and rollbacks?**
```bash
# Rolling update
kubectl set image deployment/model-server model=my-model:v3

# Check rollout status
kubectl rollout status deployment/model-server

# Rollback if issues
kubectl rollout undo deployment/model-server
```
- Strategy: `RollingUpdate` (default) or `Recreate`
- Configure `maxSurge` and `maxUnavailable` for zero-downtime deploys

**Q5: Explain ConfigMaps and Secrets.**
```yaml
# ConfigMap for non-sensitive config
apiVersion: v1
kind: ConfigMap
data:
  MODEL_NAME: "gpt-4"
  MAX_TOKENS: "4096"

# Secret for sensitive data (base64 encoded)
apiVersion: v1
kind: Secret
data:
  API_KEY: YWJjMTIz  # base64 encoded
```

### Advanced

**Q6: How do you serve ML models on Kubernetes at scale?**
- **KServe** (formerly KFServing): standard ML model serving on K8s
- Features: autoscaling (including scale-to-zero), canary rollouts, multi-model serving
- **Triton Inference Server**: NVIDIA's high-performance serving (dynamic batching, GPU sharing)
- **Seldon Core**: advanced ML deployment with A/B testing, outlier detection

---

## 🔌 APIs & Backend Services

### Beginner

**Q1: What is REST API?**
- Representational State Transfer — architectural style for web APIs
- **HTTP Methods**: GET (read), POST (create), PUT (update), DELETE (remove)
- **Stateless**: each request contains all needed information

```python
from fastapi import FastAPI
app = FastAPI()

@app.post("/predict")
async def predict(text: str):
    embedding = model.encode(text)
    return {"embedding": embedding.tolist()}
```

**Q2: What are HTTP status codes?**

| Code | Meaning | When |
|------|---------|------|
| 200 | OK | Successful request |
| 201 | Created | Resource created |
| 400 | Bad Request | Invalid input |
| 401 | Unauthorized | No/invalid auth |
| 403 | Forbidden | No permission |
| 404 | Not Found | Resource missing |
| 429 | Too Many Requests | Rate limited |
| 500 | Internal Server Error | Server bug |

**Q3: What is the difference between REST and gRPC?**

| Feature | REST | gRPC |
|---------|------|------|
| Protocol | HTTP/1.1 or 2 | HTTP/2 |
| Format | JSON (text) | Protobuf (binary) |
| Speed | Slower | 2-10x faster |
| Streaming | Limited | Bidirectional |
| Use case | Public APIs | Internal microservices |

### Intermediate

**Q4: Explain authentication and authorization.**
- **Authentication** (AuthN): WHO are you? → API keys, JWT, OAuth2
- **Authorization** (AuthZ): WHAT can you do? → RBAC, ABAC
```python
from fastapi import Depends, HTTPException
from fastapi.security import HTTPBearer

security = HTTPBearer()

@app.post("/predict")
async def predict(token = Depends(security)):
    user = verify_jwt(token.credentials)
    if "predict" not in user.permissions:
        raise HTTPException(403, "Insufficient permissions")
    return {"result": model.predict(data)}
```

**Q5: What is rate limiting and why is it important?**
- Controls request frequency per client → prevents abuse, ensures fair usage
- Algorithms: **Token Bucket**, Sliding Window, Fixed Window
- Implementation: Redis-based counters, API Gateway rules
- Critical for: LLM APIs (expensive), public endpoints, shared resources

**Q6: Explain caching strategies.**
- **Cache-aside**: app checks cache first, falls back to DB
- **Write-through**: write to cache and DB simultaneously
- **TTL**: time-based expiration
- **For AI**: cache embeddings, frequent query results, LLM responses (deterministic)
```python
import redis
cache = redis.Redis()

def get_embedding(text):
    key = hashlib.sha256(text.encode()).hexdigest()
    cached = cache.get(key)
    if cached:
        return json.loads(cached)
    embedding = model.encode(text)
    cache.setex(key, 3600, json.dumps(embedding.tolist()))
    return embedding
```

### Advanced

**Q7: Explain microservices architecture for AI systems.**
```
API Gateway → Load Balancer
  ├── Auth Service
  ├── Embedding Service (GPU)
  ├── Retrieval Service (Vector DB)
  ├── LLM Service (generation)
  ├── Monitoring Service
  └── Caching Layer (Redis)
```
- Each service independently scalable and deployable
- Communication: REST/gRPC (sync), Message Queues (async)
- Service discovery: Kubernetes DNS, Consul

---

## 🌐 Distributed Systems

### Beginner

**Q1: What is a distributed system?**
- Multiple computers working together as a single system
- **Why**: scale beyond one machine, fault tolerance, geographic distribution
- **Challenges**: network failures, clock sync, consistency, partition tolerance

**Q2: Explain CAP theorem.**
- **Consistency**: every read gets the most recent write
- **Availability**: every request gets a response
- **Partition tolerance**: system works despite network splits
- **You can only guarantee 2 out of 3** during a partition
- **CP**: consistent but may be unavailable (HBase, MongoDB)
- **AP**: available but may be inconsistent (Cassandra, DynamoDB)

**Q3: What is load balancing?**
- Distribute requests across multiple servers
- Algorithms: Round Robin, Least Connections, Weighted, IP Hash
- Types: L4 (TCP) vs L7 (HTTP/application-level)
- Tools: Nginx, HAProxy, Cloud Load Balancers

### Intermediate

**Q4: Explain message queues and event-driven architecture.**
```
Producer → [Message Queue] → Consumer

Use Cases for AI:
  - Async model inference (batch processing)
  - Data pipeline events
  - Decoupling services
```
- Tools: **Kafka** (high throughput), RabbitMQ (flexible routing), Redis Streams, SQS
- Patterns: pub/sub, work queues, event sourcing

**Q5: What is database sharding?**
- Split data across multiple databases by a shard key
- **Horizontal**: rows distributed across shards
- **Strategies**: hash-based, range-based, geographic
- **Challenges**: cross-shard queries, rebalancing, joins

**Q6: Explain eventual consistency.**
- After an update, replicas will **eventually** converge to the same state
- Tradeoff for higher availability and lower latency
- Common in: NoSQL databases, DNS, CDNs
- AI context: embedding index updates, model cache invalidation

### Advanced

**Q7: How would you design a scalable AI inference system?**
```
                    ┌─────────────────┐
  Client ──→ API Gateway ──→ Load Balancer
                    │
        ┌───────────┼───────────┐
        ▼           ▼           ▼
   GPU Worker 1  GPU Worker 2  GPU Worker N
        │           │           │
        └───────────┼───────────┘
                    ▼
              Model Registry
                    │
        ┌───────────┼───────────┐
        ▼           ▼           ▼
     Redis       Vector DB   Object Store
    (cache)     (retrieval)   (models)
```
Key decisions:
- **Batching**: dynamic batching for GPU utilization
- **Queuing**: async requests during traffic spikes
- **Autoscaling**: GPU-aware scaling (custom metrics)
- **Caching**: embedding cache, response cache
- **Health checks**: model liveness, GPU memory monitoring

**Q8: Explain consensus algorithms (Raft/Paxos).**
- Ensure multiple nodes agree on a value despite failures
- **Raft**: leader-based, easier to understand
  - Leader election → log replication → safety
  - Used in: etcd (Kubernetes), CockroachDB
- **Paxos**: more theoretical, harder to implement
- Needed for: distributed model registries, config management

**Q9: What is the circuit breaker pattern?**
```
CLOSED (normal) → errors exceed threshold → OPEN (fail fast)
    → after timeout → HALF-OPEN (test with limited traffic)
    → success → CLOSED / failure → OPEN
```
- Prevents cascading failures in microservices
- Critical for AI: LLM API failures, embedding service outages
- Implement with: resilience4j (Java), tenacity (Python), Polly (.NET)
