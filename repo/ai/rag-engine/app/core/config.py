from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    milvus_host: str = "localhost"
    milvus_port: int = 19530
    embedding_model: str = "BAAI/bge-m3"
    embedding_dim: int = 1024
    # Internal ANI Gateway address for token validation
    ani_gateway_url: str = "http://ani-gateway.ani-system.svc.cluster.local:8080"
    # HuggingFace mirror for offline/domestic environments
    hf_endpoint: str = "https://hf-mirror.com"

    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"


settings = Settings()
