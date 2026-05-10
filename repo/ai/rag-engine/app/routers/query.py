"""Knowledge base query (RAG) endpoint."""
from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, Field

router = APIRouter()


class QueryRequest(BaseModel):
    kb_id: str
    tenant_id: str
    question: str = Field(min_length=1, max_length=2000)
    session_id: str | None = None          # for multi-turn context
    top_k: int = Field(default=5, ge=1, le=20)
    score_threshold: float = Field(default=0.3, ge=0.0, le=1.0)
    inference_service_name: str = "default"   # which vLLM instance to use


class SourceChunk(BaseModel):
    doc_id: str
    file_name: str
    page: int | None
    content: str
    score: float


class QueryResponse(BaseModel):
    answer: str
    sources: list[SourceChunk]
    session_id: str
    input_tokens: int
    output_tokens: int


@router.post("/{kb_id}/query", response_model=QueryResponse)
async def query_kb(kb_id: str, req: QueryRequest) -> QueryResponse:
    """
    Hybrid RAG query: vector search (Milvus) + keyword search (PG pg_trgm) → RRF → LLM.

    When score_threshold is not met for any chunk, returns a structured
    "no relevant content found" response rather than hallucinating an answer.
    """
    if kb_id != req.kb_id:
        raise HTTPException(status_code=400, detail="path and body kb_id must match")

    # TODO:
    # 1. embed(req.question) → vector search in Milvus collection kb_{kb_id}
    # 2. keyword search in PostgreSQL (pg_trgm full-text)
    # 3. RRF merge and re-rank
    # 4. if max_score < score_threshold → return no-result response
    # 5. build prompt with context → call vLLM via ANI Gateway /v1/chat/completions
    # 6. parse response, attach sources, log tokens to DB

    return QueryResponse(
        answer="stub: not yet implemented",
        sources=[],
        session_id=req.session_id or "new-session",
        input_tokens=0,
        output_tokens=0,
    )
