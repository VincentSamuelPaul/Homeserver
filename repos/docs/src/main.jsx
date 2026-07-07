import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import RehoboamDocs from './RehoboamDocs.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RehoboamDocs />
  </StrictMode>
)
