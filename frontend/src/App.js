import Logo from "./constants/Logo";
import "./App.css";

import TextField from "@mui/material/TextField";

import { Checkbox } from "@mui/material";
import { Button } from "@mui/material";

function App() {
  return (
    <div className="App">
      <Logo />
      <div className="sub">
        <h1>SWIGGY</h1>
        <h2>STATS</h2>
        <p>This is a thirdparty tool not created by swiggy.</p>
      </div>

      <div className="main">
        <TextField
          className="input"
          id="outlined-basic"
          label="Mobile Number"
          variant="outlined"
        />
        <div className="consent">
          <Checkbox />
          <p>I understand that this tool is not affiliated with Swiggy.</p>
        </div>
        <Button
          sx={{
            marginTop: ".75xrem",
          }}
          variant="contained"
        >
          Enter
        </Button>
      </div>
    </div>
  );
}

export default App;
