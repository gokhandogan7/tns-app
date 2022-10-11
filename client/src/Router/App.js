/* eslint-disable react/prop-types */
/* eslint-disable react/react-in-jsx-scope */
import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { DetailPage, HomePage } from "../pages";
import { ErrorBoundary } from "react-error-boundary";
import ErrorPage from "../components/ui/errorPage";

function ErrorFallback({ error, resetErrorBoundary }) {
  return <ErrorPage error={error} reset={resetErrorBoundary} />;
}

function App() {
  return (
    <ErrorBoundary
      FallbackComponent={ErrorFallback}
      onReset={() => {
        // reset the state of your app so the error doesn't happen again
      }}
    >
      <Router>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/details" element={<DetailPage />} />
          <Route path="*" element={<ErrorPage/>}/>
        </Routes>
      </Router>
    </ErrorBoundary>
  );
}

export default App;
