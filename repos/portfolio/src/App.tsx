import './App.css'
import Main from './comp/MainHook'
import TechStack from './comp/TechStack'
import Projects from './comp/Projects'
import MileStones from './comp/MileStones'
import Footer from './comp/Footer'

// import { SmoothCursor } from "./components/ui/smooth-cursor";
// import { FlickeringGrid } from './components/magicui/flickering-grid'

function App() {

  return (
    <div className='w-full h-full'>
      <Main/>
      <TechStack/>
      <Projects/>
      <MileStones/>
      <Footer/>
    </div>
  )
}

export default App
