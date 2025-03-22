# 📦 Improving Durability In Kubernetes ETCD by Reducing Write Amplification

## 📝 Abstract

This research presents a durability-optimized implementation of **ETCD**—a distributed key-value store integral to Kubernetes—using **LedgerDB** and **BadgerDB** data structures. While LedgerDB ensures immutability and tamper-evidence with its append-only model, BadgerDB demonstrates high performance and lower write amplification through its LSM tree-based architecture.

Both implementations aim to address ETCD’s performance bottlenecks, specifically write amplification, as data scales. The proposed BadgerDB-based ETCD shows significant reductions in **write amplification** and **latency**, and improvements in **memory usage** and **CPU efficiency**.

Experimental analyses were conducted across ETCD data store sizes ranging from **16GB to 64GB**, and the performance was benchmarked using key operational timings and resource utilization metrics.

---

## 📚 Publication Details

- **Journal**: International Journal of Leading Research Publication (IJLRP)  
- **Publication Number (Paper ID)**: 1110  
- **Link**: https://www.ijlrp.com/research-paper.php?id=1110  
- **ISSN**: 2582-8010  
- **Impact Factor**: 9.56  
- **Publication Info**: Volume 4, Issue 2, February 2023  

---

## ✅ Key Contributions

- ✔️ Configured and optimized the **Kubernetes cluster environment**, and developed **LedgerDB** and **BadgerDB** structures in **Go** to analyze and enhance ETCD durability.
- ✔️ Conducted **detailed system comparisons** and presented data through **comprehensive tables** and **visual graphs**, focusing on **read/write latency**, **write amplification**, and **memory usage**.
- ✔️ Summarized key findings, proposed **future research directions**, and identified opportunities to enhance the performance and scalability of **distributed storage systems**.

---

## 🚀 Relevance and Impact

- 🚀 The optimized **cluster setup** and custom **LedgerDB/BadgerDB implementations** significantly enhance **data processing efficiency**, **durability**, and **scalability** in distributed systems.
- 🔬 Comprehensive system evaluations and **visual data presentations** offer clear evidence of **reduced write amplification** and **improved resource utilization**.
- 🌐 The conclusions and proposed future directions support ongoing innovation and contribute to the advancement of **high-performance**, **reliable distributed architectures**.

---

## 📊 Experimental Results (Summary)

| Store Size (GB) | Read Latency (ms) | Write Latency (ms) | Write Amplification | Memory Usage (MB) |
|-----------------|-------------------|--------------------|---------------------|-------------------|
| 16 GB           | 2.3               | 4.5                | 5.5                 | 100               |
| 24 GB           | 2.5               | 4.9                | 6.0                 | 120               |
| 32 GB           | 2.8               | 5.4                | 6.5                 | 140               |
| 40 GB           | 3.1               | 5.8                | 7.0                 | 160               |
| 48 GB           | 3.3               | 6.2                | 7.5                 | 180               |
| 64 GB           | 3.6               | 6.7                | 8.0                 | 200               |

> **Note**: BadgerDB consistently demonstrated superior performance across all metrics when compared to LedgerDB.

---

## 🔗 Citation

If you use this work, please cite it as follows:

Satya Ram Tsaliki, & Dr. B. Purnachandra Rao. (2023). Improving Durability In Kubernetes ETCD by Reducing Write Amplification. International Journal of Leading Research Publication (IJLRP), 4(2), 1–15. Paper ID: 1110. ISSN: 2582-8010.

```bibtex
@article{tsaliki2023etcd,
  title={Improving Durability In Kubernetes ETCD by Reducing Write Amplification},
  author={Satya Ram Tsaliki and B. Purnachandra Rao},
  journal={International Journal of Leading Research Publication (IJLRP)},
  volume={4},
  number={2},
  pages={1-15},
  year={2023},
  issn={2582-8010},
  note={Paper ID: 1110}
}
