"""Document ingestion and parsing endpoints."""
from fastapi import APIRouter, HTTPException
from pydantic import BaseModel

router = APIRouter()


class ParseRequest(BaseModel):
    kb_id: str
    doc_id: str
    tenant_id: str
    storage_path: str   # MinIO object path
    file_type: str      # pdf | docx | xlsx | txt | md
    # Idempotency: if chunk records already exist for this doc_id, skip re-parsing
    idempotency_key: str


class ParseResponse(BaseModel):
    doc_id: str
    chunk_count: int
    status: str


@router.post("/{kb_id}/documents/{doc_id}/parse", response_model=ParseResponse)
async def parse_document(kb_id: str, doc_id: str, req: ParseRequest) -> ParseResponse:
    """
    Parse a document and index its chunks into Milvus.

    Idempotent: calling with the same doc_id twice is safe.
    The vector primary key is deterministic: f"{doc_id}_{chunk_index}"
    so re-inserting the same chunk overwrites the previous record (upsert).
    """
    if kb_id != req.kb_id or doc_id != req.doc_id:
        raise HTTPException(status_code=400, detail="path and body kb_id/doc_id must match")

    # TODO: download from MinIO → parse with Docling/PaddleOCR → chunk → embed → upsert Milvus
    return ParseResponse(doc_id=doc_id, chunk_count=0, status="stub")


@router.delete("/{kb_id}/documents/{doc_id}/index")
async def delete_document_index(kb_id: str, doc_id: str) -> dict:
    """Remove all vectors for a document from Milvus. Idempotent."""
    # TODO: delete by expr f'doc_id == "{doc_id}"' in collection kb_{kb_id}
    return {"deleted": True, "doc_id": doc_id}
