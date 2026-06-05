# Part 3: LLMs, RAG, AI Agents & Vector Databases

---

## 📚 RAG (Retrieval-Augmented Generation)

### Beginner

**Q1: What is RAG and why is it needed?**
```
User Query → Retrieve relevant docs → Augment prompt with context → Generate answer
```
- Solves: hallucinations, stale knowledge, domain-specific questions
- No need to fine-tune the model for every domain
- Source attribution / citations become possible

**Q2: Explain the basic RAG pipeline.**
```
INDEXING (offline):
  Documents → Chunk → Embed → Store in Vector DB

RETRIEVAL (runtime):
  Query → Embed → Vector Search → Top-K chunks

GENERATION:
  Prompt = System + Retrieved Chunks + User Query → LLM → Answer
```

**Q3: What is chunking and why does it matter?**
- Breaking documents into smaller pieces for embedding
- **Too small**: loses context, irrelevant matches
- **Too large**: dilutes relevance, exceeds context window
- Common sizes: 256-1024 tokens with 10-20% overlap

### Intermediate

**Q4: Explain different chunking strategies.**

| Strategy | Description | Best For |
|----------|-------------|----------|
| Fixed-size | Split every N tokens with overlap | General purpose |
| Recursive | Split by separators (\\n\\n → \\n → . → space) | Structured text |
| Semantic | Split at topic/meaning boundaries | Research papers |
| Document-based | Respect natural boundaries (pages, sections) | PDFs, books |
| Agentic | Use LLM to determine chunk boundaries | High-quality RAG |

**Q5: What are the retrieval strategies beyond basic vector search?**
- **Hybrid search**: combine vector similarity + keyword (BM25) search
- **Re-ranking**: use cross-encoder to re-score top-K results
- **HyDE**: generate hypothetical document from query, then search
- **Multi-query**: rephrase query multiple ways, merge results
- **Parent-child**: embed small chunks, retrieve parent documents

**Q6: How do you evaluate RAG systems?**
- **Retrieval metrics**: Hit Rate, MRR (Mean Reciprocal Rank), NDCG, Precision@K
- **Generation metrics**: Faithfulness, Answer Relevancy, Context Precision
- **Frameworks**: RAGAS, DeepEval, TruLens
- **Human evaluation**: factual accuracy, completeness, groundedness

### Advanced

**Q7: Explain advanced RAG patterns.**
```
NAIVE RAG:
  Query → Retrieve → Generate

ADVANCED RAG:
  Query → Rewrite/Expand → Retrieve → Rerank → Filter → Generate

MODULAR RAG:
  Query → Route → [Web Search | Vector DB | SQL | API] → Merge → Generate

GRAPH RAG:
  Documents → Extract Entities & Relations → Knowledge Graph
  Query → Graph Traversal + Vector Search → Generate
```

**Q8: How do you handle multi-modal RAG?**
- **Text + Images**: embed images with CLIP/SigLIP, store alongside text embeddings
- **Tables**: extract to markdown/JSON, embed as text
- **Audio/Video**: transcribe → chunk → embed
- **Strategy**: unified embedding space or separate indexes with fusion

**Q9: What are the production challenges of RAG?**
- **Latency**: embedding + retrieval + generation can be slow → cache, pre-compute, streaming
- **Cost**: embedding API calls, LLM tokens → batch, cache, smaller models
- **Freshness**: documents update → incremental indexing, TTL on embeddings
- **Security**: PII in documents → access control per chunk, data filtering
- **Quality**: garbage in → garbage out → invest in data cleaning and chunking

---

## 🗃️ Vector Databases

### Beginner

**Q1: What is a vector database?**
- Specialized database for storing and searching high-dimensional vectors (embeddings)
- Core operation: **Approximate Nearest Neighbor (ANN)** search
- Unlike traditional DB (exact match), finds semantically similar items

**Q2: Name popular vector databases and compare them.**

| Database | Type | Key Feature |
|----------|------|-------------|
| Pinecone | Managed SaaS | Easiest to use, serverless |
| Weaviate | Open-source | Multi-modal, GraphQL API |
| Qdrant | Open-source | Rust-based, fast filtering |
| Milvus | Open-source | Highly scalable, GPU support |
| ChromaDB | Open-source | Simple, great for prototyping |
| pgvector | Extension | Postgres-native, familiar SQL |
| FAISS | Library | Facebook's in-memory ANN (not a DB) |

**Q3: What distance metrics are used?**
```
Cosine Similarity:   cos(A,B) = (A·B) / (||A|| × ||B||)    → [-1, 1]
Euclidean (L2):      d = √Σ(aᵢ - bᵢ)²                     → [0, ∞)
Dot Product:         A·B = Σ(aᵢ × bᵢ)                      → (-∞, ∞)
```
- **Cosine**: normalized, most common for text embeddings
- **Euclidean**: when magnitude matters
- **Dot Product**: when vectors are pre-normalized (equivalent to cosine)

### Intermediate

**Q4: Explain ANN indexing algorithms.**

| Algorithm | How it works | Pros | Cons |
|-----------|-------------|------|------|
| **HNSW** | Hierarchical graph with skip connections | Fast, accurate | High memory |
| **IVF** | Clusters vectors, searches relevant clusters | Memory efficient | Lower recall |
| **PQ** | Compresses vectors into sub-quantized codes | Very compact | Lossy |
| **ScaNN** | Google's quantization + reranking | State-of-art | Complex |

**Q5: What is metadata filtering?**
```python
# Search with filters
results = collection.query(
    query_embedding=embed("AI safety"),
    filter={
        "source": {"$eq": "arxiv"},
        "year": {"$gte": 2024},
        "category": {"$in": ["safety", "alignment"]}
    },
    top_k=10
)
```
- Pre-filtering: apply filters before ANN search (exact but slower)
- Post-filtering: apply after ANN search (faster but may return fewer results)

**Q6: How do you choose embedding dimensions?**
- **Small (384)**: `all-MiniLM-L6-v2` — fast, good for prototyping
- **Medium (768)**: `bge-base`, `e5-base` — balanced
- **Large (1024-3072)**: `text-embedding-3-large`, `bge-large` — highest quality
- Tradeoff: higher dims = more accurate but more storage/compute

### Advanced

**Q7: How do you scale a vector database in production?**
- **Sharding**: distribute vectors across nodes by partition key
- **Replication**: read replicas for high availability
- **Tiered storage**: hot (memory) → warm (SSD) → cold (disk)
- **Quantization**: reduce vector size (FP32 → INT8) for 4× compression
- **Index tuning**: adjust HNSW parameters (`M`, `efConstruction`, `efSearch`)

**Q8: Explain hybrid search implementation.**
```python
# Combine dense (vector) + sparse (BM25) retrieval
dense_results = vector_db.search(query_embedding, top_k=20)
sparse_results = bm25_index.search(query_text, top_k=20)

# Reciprocal Rank Fusion
def rrf(results_list, k=60):
    scores = defaultdict(float)
    for results in results_list:
        for rank, doc in enumerate(results):
            scores[doc.id] += 1.0 / (k + rank + 1)
    return sorted(scores.items(), key=lambda x: -x[1])

final_results = rrf([dense_results, sparse_results])
```

---

## 🤖 AI Agents

### Beginner

**Q1: What is an AI Agent?**
- An LLM-powered system that can **reason**, **plan**, and **take actions**
- Goes beyond simple Q&A — can use tools, interact with external systems
- Core loop: **Observe → Think → Act → Observe...**

**Q2: What are tools/functions in the context of agents?**
```python
tools = [
    {
        "name": "search_database",
        "description": "Search the product database",
        "parameters": {
            "query": {"type": "string"},
            "limit": {"type": "integer", "default": 10}
        }
    }
]
# LLM decides WHEN and HOW to call tools based on user query
```

**Q3: Name popular agent frameworks.**
- **LangChain / LangGraph**: most popular, graph-based agent workflows
- **CrewAI**: multi-agent collaboration framework
- **AutoGen**: Microsoft's multi-agent conversation framework
- **Semantic Kernel**: Microsoft's enterprise SDK
- **Haystack**: end-to-end NLP/RAG framework

### Intermediate

**Q4: Explain the ReAct pattern.**
```
Thought: I need to find the current stock price of NVIDIA
Action: search_web("NVIDIA stock price today")
Observation: NVIDIA (NVDA) is trading at $142.50
Thought: Now I have the price, I can answer the user
Action: respond("NVIDIA is currently trading at $142.50")
```
- Interleaves **reasoning** (chain-of-thought) with **acting** (tool use)
- More reliable than pure action-taking

**Q5: What is function calling in LLMs?**
```python
response = client.chat.completions.create(
    model="gpt-4",
    messages=[{"role": "user", "content": "What's the weather in NYC?"}],
    tools=[{
        "type": "function",
        "function": {
            "name": "get_weather",
            "parameters": {
                "type": "object",
                "properties": {"location": {"type": "string"}},
                "required": ["location"]
            }
        }
    }]
)
# LLM returns: {"name": "get_weather", "arguments": {"location": "NYC"}}
```

**Q6: How do you handle agent memory?**
- **Short-term**: conversation history (sliding window / token limit)
- **Long-term**: vector store for past interactions
- **Working memory**: scratchpad for current task state
- **Summary memory**: periodically summarize old conversations

### Advanced

**Q7: Explain multi-agent architectures.**
```
HIERARCHICAL:
  Orchestrator Agent → [Research Agent, Code Agent, Review Agent]
  
COLLABORATIVE:
  Agent A ↔ Agent B ↔ Agent C (peer discussion)

COMPETITIVE:
  Generator Agent → Critic Agent → Refined Output
```

**Q8: What are the key challenges in production agents?**
- **Reliability**: agents can go off-track → guardrails, max iterations, fallbacks
- **Cost**: multiple LLM calls per request → caching, smaller models for routing
- **Latency**: tool chains can be slow → parallel tool calls, streaming
- **Security**: prompt injection, tool misuse → input validation, sandboxing
- **Observability**: hard to debug → structured logging, tracing (LangSmith, Arize)

**Q9: Explain agent evaluation strategies.**
- **Task completion rate**: did the agent achieve the goal?
- **Tool use accuracy**: did it pick the right tools with correct parameters?
- **Trajectory evaluation**: was the reasoning path efficient?
- **Human evaluation**: quality, helpfulness, safety ratings
- **Benchmark suites**: GAIA, AgentBench, SWE-bench
