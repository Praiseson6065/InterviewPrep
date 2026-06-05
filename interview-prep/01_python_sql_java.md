# Part 1: Python, SQL & Java — Interview Prep

---

## 🐍 Python

### Beginner

**Q1: What are Python's mutable vs immutable types?**
- **Mutable**: `list`, `dict`, `set` — can be changed in place
- **Immutable**: `int`, `float`, `str`, `tuple`, `frozenset` — cannot be modified after creation

```python
# Immutable - creates new object
x = "hello"
x = x + " world"  # new string object

# Mutable - modifies in place
lst = [1, 2, 3]
lst.append(4)  # same object modified
```

**Q2: Explain list vs tuple vs set.**

| Feature | List | Tuple | Set |
|---------|------|-------|-----|
| Syntax | `[1,2,3]` | `(1,2,3)` | `{1,2,3}` |
| Mutable | ✅ | ❌ | ✅ |
| Ordered | ✅ | ✅ | ❌ |
| Duplicates | ✅ | ✅ | ❌ |
| Hashable | ❌ | ✅ | ❌ |

**Q3: What is a list comprehension?**
```python
squares = [x**2 for x in range(10) if x % 2 == 0]
# [0, 4, 16, 36, 64]
```

**Q4: Explain `*args` and `**kwargs`.**
```python
def func(*args, **kwargs):
    # args = tuple of positional arguments
    # kwargs = dict of keyword arguments
    print(args)    # (1, 2, 3)
    print(kwargs)  # {'name': 'AI'}

func(1, 2, 3, name="AI")
```

**Q5: What is the difference between `==` and `is`?**
- `==` checks **value equality**
- `is` checks **identity** (same object in memory)

```python
a = [1, 2]; b = [1, 2]
a == b  # True (same value)
a is b  # False (different objects)
```

### Intermediate

**Q6: Explain decorators with an example.**
```python
import functools, time

def timer(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        print(f"{func.__name__} took {time.time()-start:.2f}s")
        return result
    return wrapper

@timer
def train_model(epochs):
    time.sleep(epochs * 0.1)
```

**Q7: What are generators and why use them?**
```python
def read_large_file(path):
    with open(path) as f:
        for line in f:
            yield line.strip()

# Memory efficient - processes one line at a time
for line in read_large_file("data.csv"):
    process(line)
```
> Generators use **lazy evaluation** — ideal for large datasets, streaming data, or infinite sequences.

**Q8: Explain Python's GIL (Global Interpreter Lock).**
- The GIL allows only **one thread** to execute Python bytecode at a time
- **CPU-bound** tasks: use `multiprocessing` instead of `threading`
- **I/O-bound** tasks: `threading` or `asyncio` work fine (GIL is released during I/O)

**Q9: What is the difference between `deepcopy` and `copy`?**
```python
import copy
original = [[1, 2], [3, 4]]
shallow = copy.copy(original)      # nested lists still shared
deep = copy.deepcopy(original)     # completely independent
```

**Q10: Explain context managers.**
```python
class DatabaseConnection:
    def __enter__(self):
        self.conn = create_connection()
        return self.conn
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        self.conn.close()
        return False  # don't suppress exceptions

with DatabaseConnection() as conn:
    conn.execute("SELECT * FROM users")
```

### Advanced

**Q11: Explain Python's MRO (Method Resolution Order).**
```python
class A: pass
class B(A): pass
class C(A): pass
class D(B, C): pass

print(D.__mro__)
# D -> B -> C -> A -> object (C3 linearization)
```

**Q12: What are metaclasses?**
```python
class SingletonMeta(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super().__call__(*args, **kwargs)
        return cls._instances[cls]

class Database(metaclass=SingletonMeta):
    pass
```
> Metaclasses control **class creation** — used for ORMs, validation frameworks, and singleton patterns.

**Q13: Explain `asyncio` and when to use it.**
```python
import asyncio, aiohttp

async def fetch_embeddings(texts: list[str]):
    async with aiohttp.ClientSession() as session:
        tasks = [call_api(session, t) for t in texts]
        return await asyncio.gather(*tasks)

async def call_api(session, text):
    async with session.post(URL, json={"input": text}) as resp:
        return await resp.json()
```
> Use for **I/O-bound concurrency**: API calls, DB queries, file I/O. Not for CPU-heavy ML training.

**Q14: How does Python memory management work?**
- **Reference counting** as primary mechanism
- **Garbage collector** for cyclic references (generational GC)
- **Memory pools** (`pymalloc`) for small objects
- `__slots__` to reduce per-instance memory in classes

**Q15: Explain descriptor protocol.**
```python
class Validated:
    def __set_name__(self, owner, name):
        self.name = name
    def __get__(self, obj, objtype=None):
        return getattr(obj, f'_{self.name}', None)
    def __set__(self, obj, value):
        if not isinstance(value, (int, float)):
            raise TypeError(f"{self.name} must be numeric")
        setattr(obj, f'_{self.name}', value)

class Model:
    learning_rate = Validated()
    epochs = Validated()
```

---

## 🗄️ SQL

### Beginner

**Q1: What is the difference between WHERE and HAVING?**
```sql
-- WHERE filters rows BEFORE grouping
SELECT department, COUNT(*) FROM employees
WHERE salary > 50000
GROUP BY department;

-- HAVING filters groups AFTER grouping
SELECT department, AVG(salary) as avg_sal FROM employees
GROUP BY department
HAVING AVG(salary) > 60000;
```

**Q2: Explain JOIN types with examples.**
```sql
-- INNER JOIN: only matching rows
SELECT e.name, d.dept_name
FROM employees e INNER JOIN departments d ON e.dept_id = d.id;

-- LEFT JOIN: all from left + matching from right
SELECT e.name, d.dept_name
FROM employees e LEFT JOIN departments d ON e.dept_id = d.id;

-- FULL OUTER JOIN: all from both tables
-- CROSS JOIN: cartesian product
```

**Q3: What is the difference between UNION and UNION ALL?**
- `UNION` removes duplicates (slower)
- `UNION ALL` keeps all rows (faster)

**Q4: Explain GROUP BY and aggregate functions.**
```sql
SELECT category, 
       COUNT(*) as total,
       AVG(price) as avg_price,
       MAX(price) as max_price
FROM products
GROUP BY category
ORDER BY avg_price DESC;
```

### Intermediate

**Q5: Explain window functions.**
```sql
SELECT employee_name, department, salary,
    RANK() OVER (PARTITION BY department ORDER BY salary DESC) as dept_rank,
    LAG(salary) OVER (ORDER BY salary) as prev_salary,
    SUM(salary) OVER (PARTITION BY department) as dept_total,
    AVG(salary) OVER (
        ORDER BY hire_date ROWS BETWEEN 2 PRECEDING AND CURRENT ROW
    ) as moving_avg
FROM employees;
```

**Q6: What are indexes and when to use them?**
```sql
-- B-tree index (default) - good for equality & range
CREATE INDEX idx_salary ON employees(salary);

-- Composite index - column order matters
CREATE INDEX idx_dept_sal ON employees(department, salary);

-- Partial index
CREATE INDEX idx_active ON employees(email) WHERE is_active = true;
```
> **When NOT to index**: small tables, frequently updated columns, low-cardinality columns.

**Q7: Explain CTEs (Common Table Expressions).**
```sql
WITH monthly_revenue AS (
    SELECT DATE_TRUNC('month', order_date) as month,
           SUM(amount) as revenue
    FROM orders
    GROUP BY 1
),
growth AS (
    SELECT month, revenue,
           LAG(revenue) OVER (ORDER BY month) as prev_revenue,
           (revenue - LAG(revenue) OVER (ORDER BY month)) / 
            LAG(revenue) OVER (ORDER BY month) * 100 as growth_pct
    FROM monthly_revenue
)
SELECT * FROM growth WHERE growth_pct > 10;
```

**Q8: What is normalization? Explain 1NF, 2NF, 3NF.**
- **1NF**: Atomic values, no repeating groups
- **2NF**: 1NF + no partial dependencies (every non-key depends on full primary key)
- **3NF**: 2NF + no transitive dependencies (non-key depends only on key)

### Advanced

**Q9: Explain query execution plans and optimization.**
```sql
EXPLAIN ANALYZE
SELECT e.name, d.dept_name
FROM employees e
JOIN departments d ON e.dept_id = d.id
WHERE e.salary > 100000;
```
Key things to look for: **Seq Scan** (missing index), **Nested Loop** vs **Hash Join**, **Sort** cost, row estimates vs actuals.

**Q10: Explain transaction isolation levels.**

| Level | Dirty Read | Non-repeatable Read | Phantom Read |
|-------|-----------|-------------------|-------------|
| READ UNCOMMITTED | ✅ | ✅ | ✅ |
| READ COMMITTED | ❌ | ✅ | ✅ |
| REPEATABLE READ | ❌ | ❌ | ✅ |
| SERIALIZABLE | ❌ | ❌ | ❌ |

**Q11: Write a recursive CTE.**
```sql
-- Org hierarchy traversal
WITH RECURSIVE org_tree AS (
    SELECT id, name, manager_id, 1 as level
    FROM employees WHERE manager_id IS NULL
    UNION ALL
    SELECT e.id, e.name, e.manager_id, t.level + 1
    FROM employees e
    JOIN org_tree t ON e.manager_id = t.id
)
SELECT * FROM org_tree ORDER BY level;
```

---

## ☕ Java

### Beginner

**Q1: Explain OOP principles in Java.**
- **Encapsulation**: private fields + public getters/setters
- **Inheritance**: `extends` for classes, `implements` for interfaces
- **Polymorphism**: method overriding (runtime) + overloading (compile-time)
- **Abstraction**: abstract classes and interfaces

**Q2: What is the difference between `==` and `.equals()`?**
```java
String a = new String("hello");
String b = new String("hello");
a == b       // false (different references)
a.equals(b)  // true (same value)
```

**Q3: Explain exception handling.**
```java
try {
    riskyOperation();
} catch (IOException e) {
    logger.error("IO failed", e);
} catch (Exception e) {
    logger.error("Unexpected", e);
} finally {
    cleanup(); // always runs
}
```
- **Checked**: must handle (`IOException`, `SQLException`)
- **Unchecked**: runtime (`NullPointerException`, `ArrayIndexOutOfBoundsException`)

### Intermediate

**Q4: Explain Java Collections Framework.**

| Interface | Implementations | Use Case |
|-----------|----------------|----------|
| `List` | `ArrayList`, `LinkedList` | Ordered, indexed |
| `Set` | `HashSet`, `TreeSet` | Unique elements |
| `Map` | `HashMap`, `TreeMap`, `ConcurrentHashMap` | Key-value pairs |
| `Queue` | `PriorityQueue`, `ArrayDeque` | FIFO/priority processing |

**Q5: What are Java Streams?**
```java
List<String> names = employees.stream()
    .filter(e -> e.getSalary() > 50000)
    .sorted(Comparator.comparing(Employee::getName))
    .map(Employee::getName)
    .collect(Collectors.toList());
```

**Q6: Explain multithreading in Java.**
```java
// ExecutorService approach (preferred)
ExecutorService executor = Executors.newFixedThreadPool(4);
Future<Result> future = executor.submit(() -> computeEmbeddings(data));
Result result = future.get(); // blocks until done

// CompletableFuture for async chaining
CompletableFuture.supplyAsync(() -> fetchData())
    .thenApply(data -> process(data))
    .thenAccept(result -> save(result));
```

### Advanced

**Q7: Explain JVM memory model.**
- **Heap**: objects (Young Gen → Old Gen)
- **Stack**: method frames, local variables
- **Metaspace**: class metadata (replaced PermGen in Java 8+)
- **GC Types**: G1 (default), ZGC (low latency), Shenandoah

**Q8: What is the `volatile` keyword vs `synchronized`?**
- `volatile`: guarantees visibility across threads (no caching), no atomicity
- `synchronized`: guarantees both visibility and atomicity (mutual exclusion)

**Q9: Explain design patterns relevant to AI systems.**
```java
// Strategy Pattern - swappable ML models
interface ModelStrategy {
    Prediction predict(Features input);
}
class RandomForestStrategy implements ModelStrategy { ... }
class NeuralNetStrategy implements ModelStrategy { ... }

// Builder Pattern - complex config
ModelConfig config = ModelConfig.builder()
    .modelName("gpt-4")
    .temperature(0.7)
    .maxTokens(4096)
    .build();

// Observer Pattern - model monitoring
interface ModelObserver {
    void onPrediction(PredictionEvent event);
}
```
