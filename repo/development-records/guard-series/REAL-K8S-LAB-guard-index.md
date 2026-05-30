# M1-REAL-LAB Guard Series Index

> Sprint 5 / REAL-K8S-LAB-A 系列 guard validators 完整索引
>
> **更新方式**：每个 guard 微批次完成后**只在本文件追加一行**（批次 ID + 类别 + guard 名称即为完整记录），不再单独建 `{批次名}.md`，不更新 README.md、CURRENT-SPRINT.md 或 ANI-06 Section 零。guard 的可验证证据由对应校验脚本（如 `scripts/validate_real_k8s_profile.py`）及其测试承载。
>
> **最新 guard ID**：`M1-REAL-LAB-KX`（KMS/SM4 live gate command type guard，2026-05-29）
>
> **Guard 数量**：299 个（B 至 KX）
>
> **历史说明**：2026-05-30 起，原先每个 guard 对应的独立 `m1-real-lab-*.md` 批次记录已合并进本索引并删除，以遏制 development-records 目录膨胀；如需追溯实现，见对应校验脚本与其测试。

---

## 类别说明

| 类别 | 含义 |
|---|---|
| `infra` | contract gate index、live evidence JSON 输出、live runner/聚合 |
| `env` | required env、env file 加载、preflight-only、gate selector、env 值守卫 |
| `summary-report` | component summary/report 完整性、provenance、whitespace 验证 |
| `evidence` | per-gate evidence JSON 内容、identity 验证 |
| `live-check-profile` | live_checks profile 字段类型/形状验证 |
| `contract-gate-profile` | contract_gates manifest/validator identity 验证 |
| `path` | 输入路径 whitespace/empty/writability 守卫 |
| `kubeconfig` | kubeconfig 字段类型、值、文件存在性验证 |
| `live-gate-cmd` | 各组件 live gate YAML command 类型守卫 |

---

## 完整 Guard 列表

| 批次 ID | 类别 | Guard 名称 |
|---|---|---|
| M1-REAL-LAB-B | `infra` | component contract gate index |
| M1-REAL-LAB-C | `infra` | `--live` evidence JSON output |
| M1-REAL-LAB-D | `infra` | component live runner |
| M1-REAL-LAB-E | `infra` | component live 聚合失败摘要 |
| M1-REAL-LAB-F | `env` | required env preflight |
| M1-REAL-LAB-G | `env` | component env template |
| M1-REAL-LAB-H | `env` | component env file loader |
| M1-REAL-LAB-I | `env` | component preflight-only mode |
| M1-REAL-LAB-J | `env` | component gate selector |
| M1-REAL-LAB-K | `summary-report` | component summary report |
| M1-REAL-LAB-L | `summary-report` | component report stale summary guard |
| M1-REAL-LAB-M | `summary-report` | component report diagnostic details |
| M1-REAL-LAB-N | `evidence` | component evidence integrity guard |
| M1-REAL-LAB-O | `evidence` | component evidence content guard |
| M1-REAL-LAB-P | `evidence` | component report passed-evidence audit |
| M1-REAL-LAB-Q | `summary-report` | component report unresolved exit guard |
| M1-REAL-LAB-R | `summary-report` | component report overall status |
| M1-REAL-LAB-S | `summary-report` | component diagnostic redaction |
| M1-REAL-LAB-T | `summary-report` | component record paths |
| M1-REAL-LAB-U | `env` | component required env name guard |
| M1-REAL-LAB-V | `infra` | component gate id uniqueness guard |
| M1-REAL-LAB-W | `summary-report` | component summary profile guard |
| M1-REAL-LAB-X | `summary-report` | component summary count guard |
| M1-REAL-LAB-Y | `summary-report` | component summary duplicate gate guard |
| M1-REAL-LAB-Z | `summary-report` | component summary gate profile guard |
| M1-REAL-LAB-AA | `evidence` | component evidence profile guard |
| M1-REAL-LAB-AB | `evidence` | component evidence gate id guard |
| M1-REAL-LAB-AC | `evidence` | component evidence identity required guard |
| M1-REAL-LAB-AD | `evidence` | component report evidence_output required guard |
| M1-REAL-LAB-AE | `summary-report` | component report summary profile required guard |
| M1-REAL-LAB-AF | `evidence` | component report unreadable evidence guard |
| M1-REAL-LAB-AG | `summary-report` | component summary status guard |
| M1-REAL-LAB-AH | `summary-report` | component report unreadable summary guard |
| M1-REAL-LAB-AI | `env` | component env file unreadable guard |
| M1-REAL-LAB-AJ | `env` | component env duplicate assignment guard |
| M1-REAL-LAB-AK | `env` | component env unknown assignment guard |
| M1-REAL-LAB-AL | `env` | component report env file validation |
| M1-REAL-LAB-AM | `path` | component evidence dir guard |
| M1-REAL-LAB-AN | `path` | component validator evidence output guard |
| M1-REAL-LAB-AO | `path` | profile output guard |
| M1-REAL-LAB-AP | `path` | profile path diagnostic guard |
| M1-REAL-LAB-AQ | `path` | profile unreadable guard |
| M1-REAL-LAB-AR | `path` | profile malformed guard |
| M1-REAL-LAB-AS | `live-check-profile` | live JSON root guard |
| M1-REAL-LAB-AT | `live-check-profile` | live JSON items guard |
| M1-REAL-LAB-AU | `live-check-profile` | live node conditions guard |
| M1-REAL-LAB-AV | `live-check-profile` | live metadata guard |
| M1-REAL-LAB-AW | `live-check-profile` | profile minimum nodes guard |
| M1-REAL-LAB-AX | `live-check-profile` | live stdout command guard |
| M1-REAL-LAB-AY | `live-check-profile` | live pass condition guard |
| M1-REAL-LAB-AZ | `live-check-profile` | live command type guard |
| M1-REAL-LAB-BA | `live-check-profile` | live command args guard |
| M1-REAL-LAB-BB | `live-check-profile` | live check field type guard |
| M1-REAL-LAB-BC | `live-check-profile` | live pass condition profile guard |
| M1-REAL-LAB-BD | `live-check-profile` | live check id uniqueness guard |
| M1-REAL-LAB-BE | `live-check-profile` | live check command profile guard |
| M1-REAL-LAB-BF | `live-check-profile` | live check JSON output profile guard |
| M1-REAL-LAB-BG | `live-check-profile` | live check safe verb profile guard |
| M1-REAL-LAB-BH | `contract-gate-profile` | contract gate command validator profile guard |
| M1-REAL-LAB-BI | `contract-gate-profile` | contract gate manifest profile guard |
| M1-REAL-LAB-BJ | `contract-gate-profile` | contract gate validator identity profile guard |
| M1-REAL-LAB-BK | `contract-gate-profile` | contract gate manifest live check id guard |
| M1-REAL-LAB-BL | `contract-gate-profile` | contract gate manifest live check field guard |
| M1-REAL-LAB-BM | `contract-gate-profile` | contract gate manifest pass condition guard |
| M1-REAL-LAB-BN | `contract-gate-profile` | contract gate manifest required check guard |
| M1-REAL-LAB-BO | `contract-gate-profile` | contract gate validator required env guard |
| M1-REAL-LAB-BP | `summary-report` | component live command audit |
| M1-REAL-LAB-BQ | `summary-report` | component report command audit guard |
| M1-REAL-LAB-BR | `summary-report` | component report command shape guard |
| M1-REAL-LAB-BS | `summary-report` | component summary passed flag guard |
| M1-REAL-LAB-BT | `summary-report` | component summary passed required guard |
| M1-REAL-LAB-BU | `summary-report` | component summary object required guard |
| M1-REAL-LAB-BV | `summary-report` | component summary count fields required guard |
| M1-REAL-LAB-BW | `summary-report` | component summary count type guard |
| M1-REAL-LAB-BX | `summary-report` | component summary gate passed type guard |
| M1-REAL-LAB-BY | `summary-report` | component summary missing_env shape guard |
| M1-REAL-LAB-BZ | `summary-report` | component summary returncode type guard |
| M1-REAL-LAB-CA | `summary-report` | component summary error type guard |
| M1-REAL-LAB-CB | `summary-report` | component summary evidence_output type guard |
| M1-REAL-LAB-CC | `summary-report` | component summary gate profile type guard |
| M1-REAL-LAB-CD | `summary-report` | component summary gate id type guard |
| M1-REAL-LAB-CE | `summary-report` | component summary top-level profile type guard |
| M1-REAL-LAB-CF | `summary-report` | component summary top-level status type guard |
| M1-REAL-LAB-CG | `summary-report` | component summary component_gates required guard |
| M1-REAL-LAB-CH | `summary-report` | component summary gate missing_env required guard |
| M1-REAL-LAB-CI | `summary-report` | component summary gate profile required guard |
| M1-REAL-LAB-CJ | `summary-report` | component summary gate required_env required guard |
| M1-REAL-LAB-CK | `summary-report` | component summary missing_env declared guard |
| M1-REAL-LAB-CL | `summary-report` | component summary status outcome guard |
| M1-REAL-LAB-CM | `summary-report` | component summary live status blocked gate guard |
| M1-REAL-LAB-CN | `summary-report` | component summary preflight status failed gate guard |
| M1-REAL-LAB-CO | `summary-report` | component summary preflight evidence audit split |
| M1-REAL-LAB-CP | `summary-report` | component live blocked preflight summary reportability |
| M1-REAL-LAB-CQ | `summary-report` | component report audited summary counts |
| M1-REAL-LAB-CR | `evidence` | component evidence identity type guard |
| M1-REAL-LAB-CS | `evidence` | component evidence status consistency guard |
| M1-REAL-LAB-CT | `evidence` | component evidence status type guard |
| M1-REAL-LAB-CU | `evidence` | component evidence passed type guard |
| M1-REAL-LAB-CV | `evidence` | component evidence status/passed mismatch diagnostic |
| M1-REAL-LAB-CW | `evidence` | component evidence missing outcome diagnostic |
| M1-REAL-LAB-CX | `evidence` | component evidence empty status diagnostic |
| M1-REAL-LAB-CY | `evidence` | component evidence non-passing status diagnostic |
| M1-REAL-LAB-CZ | `evidence` | component evidence error redaction |
| M1-REAL-LAB-DA | `evidence` | component report evidence dir binding guard |
| M1-REAL-LAB-DB | `summary-report` | component report failed gate command provenance |
| M1-REAL-LAB-DC | `summary-report` | component report validator script provenance |
| M1-REAL-LAB-DD | `summary-report` | component report passed returncode guard |
| M1-REAL-LAB-DE | `summary-report` | component report passed error guard |
| M1-REAL-LAB-DF | `summary-report` | component report passed returncode required guard |
| M1-REAL-LAB-DG | `summary-report` | component summary required_env provenance guard |
| M1-REAL-LAB-DH | `summary-report` | component summary non-empty gates guard |
| M1-REAL-LAB-DI | `summary-report` | component report validator_script exact provenance guard |
| M1-REAL-LAB-DJ | `summary-report` | component report evidence_output exact provenance guard |
| M1-REAL-LAB-DK | `summary-report` | component summary status exact provenance guard |
| M1-REAL-LAB-DL | `summary-report` | component summary profile exact provenance guard |
| M1-REAL-LAB-DM | `summary-report` | component summary gate profile exact provenance guard |
| M1-REAL-LAB-DN | `evidence` | component evidence profile exact provenance guard |
| M1-REAL-LAB-DO | `evidence` | component evidence gate id exact provenance guard |
| M1-REAL-LAB-DP | `summary-report` | component summary gate id exact provenance guard |
| M1-REAL-LAB-DQ | `summary-report` | component summary required_env exact provenance guard |
| M1-REAL-LAB-DR | `summary-report` | component summary missing_env exact provenance guard |
| M1-REAL-LAB-DS | `summary-report` | component report command entry exact provenance guard |
| M1-REAL-LAB-DT | `summary-report` | component report validator_script whitespace diagnostic |
| M1-REAL-LAB-DU | `summary-report` | component report validator_script type guard |
| M1-REAL-LAB-DV | `summary-report` | component report failed live gate returncode required guard |
| M1-REAL-LAB-DW | `summary-report` | component report failed live gate error field required guard |
| M1-REAL-LAB-DX | `summary-report` | component report failed live gate error whitespace guard |
| M1-REAL-LAB-DY | `summary-report` | component report preflight blocked gate execution metadata guard |
| M1-REAL-LAB-DZ | `summary-report` | component report preflight passed gate execution metadata guard |
| M1-REAL-LAB-EA | `summary-report` | component report preflight passed live-only provenance guard |
| M1-REAL-LAB-EB | `summary-report` | component report preflight blocked planned provenance guard |
| M1-REAL-LAB-EC | `summary-report` | component report preflight passed planned provenance guard |
| M1-REAL-LAB-ED | `summary-report` | component summary passed missing_env guard |
| M1-REAL-LAB-EE | `summary-report` | component report failed live gate error empty guard |
| M1-REAL-LAB-EF | `env` | component env value whitespace guard |
| M1-REAL-LAB-EG | `env` | component required env runtime whitespace guard |
| M1-REAL-LAB-EH | `env` | component env assignment whitespace guard |
| M1-REAL-LAB-EI | `env` | component validator live config whitespace guard |
| M1-REAL-LAB-EJ | `env` | component validator evidence output whitespace guard |
| M1-REAL-LAB-EK | `path` | profile output whitespace guard |
| M1-REAL-LAB-EL | `path` | component evidence dir whitespace guard |
| M1-REAL-LAB-EM | `path` | component env file path whitespace guard |
| M1-REAL-LAB-EN | `path` | component gate selector whitespace guard |
| M1-REAL-LAB-EO | `path` | component summary path whitespace guard |
| M1-REAL-LAB-EP | `path` | profile path whitespace guard |
| M1-REAL-LAB-EQ | `path` | kubeconfig path whitespace guard |
| M1-REAL-LAB-ER | `infra` | live evidence outcome summary |
| M1-REAL-LAB-ES | `path` | kubeconfig empty path guard |
| M1-REAL-LAB-ET | `path` | component evidence dir empty path guard |
| M1-REAL-LAB-EU | `path` | profile output empty path guard |
| M1-REAL-LAB-EV | `path` | component report empty summary path guard |
| M1-REAL-LAB-EW | `path` | component env file empty path guard |
| M1-REAL-LAB-EX | `path` | profile empty path guard |
| M1-REAL-LAB-EY | `path` | direct live evidence output empty path guard |
| M1-REAL-LAB-EZ | `path` | direct live gate path empty guard |
| M1-REAL-LAB-FA | `path` | direct live gate path whitespace guard |
| M1-REAL-LAB-FB | `path` | direct live gate path diagnostic guard |
| M1-REAL-LAB-FC | `path` | direct live gate path unreadable guard |
| M1-REAL-LAB-FD | `path` | direct live gate malformed guard |
| M1-REAL-LAB-FE | `path` | direct live doc missing guard |
| M1-REAL-LAB-FF | `path` | direct live doc decode guard |
| M1-REAL-LAB-FG | `path` | direct live evidence output directory guard |
| M1-REAL-LAB-FH | `path` | direct live evidence output parent guard |
| M1-REAL-LAB-FI | `path` | direct live evidence output parent creation guard |
| M1-REAL-LAB-FJ | `path` | direct live evidence output writability guard |
| M1-REAL-LAB-FK | `path` | direct live evidence output parent writability guard |
| M1-REAL-LAB-FL | `path` | component evidence dir writability guard |
| M1-REAL-LAB-FM | `path` | profile output writability guard |
| M1-REAL-LAB-FN | `path` | profile output parent creation diagnostic guard |
| M1-REAL-LAB-FO | `path` | profile output parent creation writability diagnostic guard |
| M1-REAL-LAB-FP | `kubeconfig` | kubeconfig missing path guard |
| M1-REAL-LAB-FQ | `kubeconfig` | kubeconfig directory path guard |
| M1-REAL-LAB-FR | `kubeconfig` | kubeconfig unreadable file guard |
| M1-REAL-LAB-FS | `kubeconfig` | kubeconfig malformed YAML guard |
| M1-REAL-LAB-FT | `kubeconfig` | kubeconfig required fields guard |
| M1-REAL-LAB-FU | `kubeconfig` | kubeconfig empty collections guard |
| M1-REAL-LAB-FV | `kubeconfig` | kubeconfig current-context reference guard |
| M1-REAL-LAB-FW | `kubeconfig` | kubeconfig context reference guard |
| M1-REAL-LAB-FX | `kubeconfig` | kubeconfig cluster server guard |
| M1-REAL-LAB-FY | `kubeconfig` | kubeconfig user object guard |
| M1-REAL-LAB-FZ | `kubeconfig` | kubeconfig user auth material guard |
| M1-REAL-LAB-GA | `kubeconfig` | kubeconfig cluster server URL guard |
| M1-REAL-LAB-GB | `kubeconfig` | kubeconfig cluster server whitespace guard |
| M1-REAL-LAB-GC | `kubeconfig` | kubeconfig user auth value guard |
| M1-REAL-LAB-GD | `kubeconfig` | kubeconfig current-context whitespace guard |
| M1-REAL-LAB-GE | `kubeconfig` | kubeconfig context cluster/user whitespace guard |
| M1-REAL-LAB-GF | `kubeconfig` | kubeconfig named entry whitespace guard |
| M1-REAL-LAB-GG | `kubeconfig` | kubeconfig collection entry object guard |
| M1-REAL-LAB-GH | `kubeconfig` | kubeconfig collection entry name required guard |
| M1-REAL-LAB-GI | `kubeconfig` | kubeconfig collection entry name uniqueness guard |
| M1-REAL-LAB-GJ | `kubeconfig` | kubeconfig collection entry payload object guard |
| M1-REAL-LAB-GK | `kubeconfig` | kubeconfig top-level identity value guard |
| M1-REAL-LAB-GL | `kubeconfig` | kubeconfig user auth value whitespace guard |
| M1-REAL-LAB-GM | `kubeconfig` | kubeconfig user auth value type guard |
| M1-REAL-LAB-GN | `kubeconfig` | kubeconfig exec args type guard |
| M1-REAL-LAB-GO | `kubeconfig` | kubeconfig exec env type guard |
| M1-REAL-LAB-GP | `kubeconfig` | kubeconfig exec env name required guard |
| M1-REAL-LAB-GQ | `kubeconfig` | kubeconfig exec env value required guard |
| M1-REAL-LAB-GR | `kubeconfig` | kubeconfig exec command required guard |
| M1-REAL-LAB-GS | `kubeconfig` | kubeconfig exec apiVersion required guard |
| M1-REAL-LAB-GT | `kubeconfig` | kubeconfig exec apiVersion supported guard |
| M1-REAL-LAB-GU | `kubeconfig` | kubeconfig exec interactiveMode supported guard |
| M1-REAL-LAB-GV | `kubeconfig` | kubeconfig exec v1 interactiveMode required guard |
| M1-REAL-LAB-GW | `kubeconfig` | kubeconfig exec provideClusterInfo type guard |
| M1-REAL-LAB-GX | `kubeconfig` | kubeconfig exec env name pattern guard |
| M1-REAL-LAB-GY | `kubeconfig` | kubeconfig exec object guard |
| M1-REAL-LAB-GZ | `kubeconfig` | kubeconfig exec command whitespace guard |
| M1-REAL-LAB-HA | `kubeconfig` | kubeconfig exec args non-empty guard |
| M1-REAL-LAB-HB | `kubeconfig` | kubeconfig exec env value non-empty guard |
| M1-REAL-LAB-HC | `kubeconfig` | kubeconfig cluster insecure-skip-tls-verify type guard |
| M1-REAL-LAB-HD | `kubeconfig` | kubeconfig cluster certificate-authority-data type guard |
| M1-REAL-LAB-HE | `kubeconfig` | kubeconfig cluster certificate-authority type guard |
| M1-REAL-LAB-HF | `kubeconfig` | kubeconfig cluster tls-server-name type guard |
| M1-REAL-LAB-HG | `kubeconfig` | kubeconfig cluster proxy-url type guard |
| M1-REAL-LAB-HH | `kubeconfig` | kubeconfig cluster disable-compression type guard |
| M1-REAL-LAB-HI | `kubeconfig` | kubeconfig user auth-provider object type guard |
| M1-REAL-LAB-HJ | `kubeconfig` | kubeconfig user auth-provider config object type guard |
| M1-REAL-LAB-HK | `kubeconfig` | kubeconfig user auth-provider name required guard |
| M1-REAL-LAB-HL | `kubeconfig` | kubeconfig user auth-provider config value type guard |
| M1-REAL-LAB-HM | `kubeconfig` | kubeconfig user auth-provider config key type guard |
| M1-REAL-LAB-HN | `kubeconfig` | kubeconfig user auth-provider config key non-empty guard |
| M1-REAL-LAB-HO | `kubeconfig` | kubeconfig user auth-provider config key whitespace guard |
| M1-REAL-LAB-HP | `kubeconfig` | kubeconfig user auth-provider config value non-empty guard |
| M1-REAL-LAB-HQ | `kubeconfig` | kubeconfig user auth-provider name whitespace guard |
| M1-REAL-LAB-HR | `kubeconfig` | kubeconfig user auth-provider config value whitespace guard |
| M1-REAL-LAB-HS | `kubeconfig` | kubeconfig cluster certificate-authority-data non-empty guard |
| M1-REAL-LAB-HT | `kubeconfig` | kubeconfig cluster certificate-authority non-empty guard |
| M1-REAL-LAB-HU | `kubeconfig` | kubeconfig cluster tls-server-name non-empty guard |
| M1-REAL-LAB-HV | `kubeconfig` | kubeconfig cluster proxy-url non-empty guard |
| M1-REAL-LAB-HW | `kubeconfig` | kubeconfig cluster proxy-url URL guard |
| M1-REAL-LAB-HX | `kubeconfig` | kubeconfig cluster proxy-url whitespace guard |
| M1-REAL-LAB-HY | `kubeconfig` | kubeconfig cluster certificate-authority-data whitespace guard |
| M1-REAL-LAB-HZ | `kubeconfig` | kubeconfig cluster certificate-authority whitespace guard |
| M1-REAL-LAB-IA | `kubeconfig` | kubeconfig cluster tls-server-name whitespace guard |
| M1-REAL-LAB-IB | `kubeconfig` | kubeconfig context namespace type guard |
| M1-REAL-LAB-IC | `kubeconfig` | kubeconfig context namespace non-empty guard |
| M1-REAL-LAB-ID | `kubeconfig` | kubeconfig context namespace whitespace guard |
| M1-REAL-LAB-IE | `kubeconfig` | kubeconfig context namespace name guard |
| M1-REAL-LAB-IF | `kubeconfig` | kubeconfig context namespace system namespace guard |
| M1-REAL-LAB-IG | `kubeconfig` | kubeconfig cluster certificate-authority file existence guard |
| M1-REAL-LAB-IH | `kubeconfig` | kubeconfig cluster certificate-authority file readable guard |
| M1-REAL-LAB-II | `kubeconfig` | kubeconfig cluster certificate-authority-data base64 guard |
| M1-REAL-LAB-IJ | `kubeconfig` | kubeconfig user tokenFile existence guard |
| M1-REAL-LAB-IK | `kubeconfig` | kubeconfig user tokenFile file type guard |
| M1-REAL-LAB-IL | `kubeconfig` | kubeconfig user tokenFile readability guard |
| M1-REAL-LAB-IM | `kubeconfig` | kubeconfig user client-certificate file existence guard |
| M1-REAL-LAB-IN | `kubeconfig` | kubeconfig user client-certificate file type guard |
| M1-REAL-LAB-IO | `kubeconfig` | kubeconfig user client-certificate readability guard |
| M1-REAL-LAB-IP | `kubeconfig` | kubeconfig user client-key file existence guard |
| M1-REAL-LAB-IQ | `kubeconfig` | kubeconfig user client-key file type guard |
| M1-REAL-LAB-IR | `kubeconfig` | kubeconfig user client-key readability guard |
| M1-REAL-LAB-IS | `kubeconfig` | kubeconfig user client-certificate-data base64 guard |
| M1-REAL-LAB-IT | `kubeconfig` | kubeconfig user client-key-data base64 guard |
| M1-REAL-LAB-IU | `kubeconfig` | kubeconfig user auth string non-empty guard |
| M1-REAL-LAB-IV | `kubeconfig` | kubeconfig exec args whitespace guard |
| M1-REAL-LAB-IW | `kubeconfig` | kubeconfig exec env name whitespace guard |
| M1-REAL-LAB-IX | `kubeconfig` | kubeconfig exec env value whitespace guard |
| M1-REAL-LAB-IY | `kubeconfig` | kubeconfig exec interactiveMode whitespace guard |
| M1-REAL-LAB-IZ | `kubeconfig` | kubeconfig exec apiVersion whitespace guard |
| M1-REAL-LAB-JA | `kubeconfig` | kubeconfig exec interactiveMode type guard |
| M1-REAL-LAB-JB | `kubeconfig` | kubeconfig exec interactiveMode empty guard |
| M1-REAL-LAB-JC | `kubeconfig` | kubeconfig exec command type guard |
| M1-REAL-LAB-JD | `kubeconfig` | kubeconfig exec apiVersion type guard |
| M1-REAL-LAB-JE | `kubeconfig` | kubeconfig exec installHint type guard |
| M1-REAL-LAB-JF | `kubeconfig` | kubeconfig exec installHint empty guard |
| M1-REAL-LAB-JG | `kubeconfig` | kubeconfig exec installHint whitespace guard |
| M1-REAL-LAB-JH | `kubeconfig` | kubeconfig exec provideClusterInfo scoped diagnostic guard |
| M1-REAL-LAB-JI | `kubeconfig` | kubeconfig exec args scoped diagnostic guard |
| M1-REAL-LAB-JJ | `kubeconfig` | kubeconfig exec env scoped diagnostic guard |
| M1-REAL-LAB-JK | `kubeconfig` | kubeconfig user auth string whitespace scoped diagnostic guard |
| M1-REAL-LAB-JL | `kubeconfig` | kubeconfig user supported auth material guard |
| M1-REAL-LAB-JM | `kubeconfig` | kubeconfig user supported auth whitespace guard |
| M1-REAL-LAB-JN | `kubeconfig` | kubeconfig user supported auth type guard |
| M1-REAL-LAB-JO | `kubeconfig` | kubeconfig exec unknown nested auth-like field guard |
| M1-REAL-LAB-JP | `kubeconfig` | kubeconfig auth-provider unknown nested auth-like field guard |
| M1-REAL-LAB-JQ | `kubeconfig` | kubeconfig supported auth object unknown nested whitespace guard |
| M1-REAL-LAB-JR | `kubeconfig` | kubeconfig supported auth object empty diagnostic guard |
| M1-REAL-LAB-JS | `kubeconfig` | kubeconfig supported auth object null diagnostic guard |
| M1-REAL-LAB-JT | `kubeconfig` | kubeconfig top-level auth string null diagnostic guard |
| M1-REAL-LAB-JU | `kubeconfig` | kubeconfig auth-provider config null diagnostic guard |
| M1-REAL-LAB-JV | `kubeconfig` | kubeconfig exec args null diagnostic guard |
| M1-REAL-LAB-JW | `kubeconfig` | kubeconfig exec env null diagnostic guard |
| M1-REAL-LAB-JX | `kubeconfig` | kubeconfig exec installHint null diagnostic guard |
| M1-REAL-LAB-JY | `kubeconfig` | kubeconfig exec provideClusterInfo null diagnostic guard |
| M1-REAL-LAB-JZ | `kubeconfig` | kubeconfig exec interactiveMode null diagnostic guard |
| M1-REAL-LAB-KA | `kubeconfig` | kubeconfig cluster insecure-skip-tls-verify null diagnostic guard |
| M1-REAL-LAB-KL | `kubeconfig` | kubeconfig cluster disable-compression null diagnostic guard |
| M1-REAL-LAB-KM | `kubeconfig` | kubeconfig cluster certificate-authority-data null diagnostic guard |
| M1-REAL-LAB-KN | `kubeconfig` | kubeconfig cluster certificate-authority null diagnostic guard |
| M1-REAL-LAB-KO | `kubeconfig` | kubeconfig cluster tls-server-name null diagnostic guard |
| M1-REAL-LAB-KP | `kubeconfig` | kubeconfig cluster proxy-url null diagnostic guard |
| M1-REAL-LAB-KQ | `kubeconfig` | kubeconfig current-context namespace null diagnostic guard |
| M1-REAL-LAB-KR | `live-gate-cmd` | vCluster live gate command type guard |
| M1-REAL-LAB-KS | `live-gate-cmd` | vCluster upgrade live gate command type guard |
| M1-REAL-LAB-KT | `live-gate-cmd` | node pool live gate command type guard |
| M1-REAL-LAB-KU | `live-gate-cmd` | Kube-OVN live gate command type guard |
| M1-REAL-LAB-KV | `live-gate-cmd` | KubeVirt live gate command type guard |
| M1-REAL-LAB-KW | `live-gate-cmd` | controller HA live gate command type guard |
| M1-REAL-LAB-KX | `live-gate-cmd` | KMS/SM4 live gate command type guard |
