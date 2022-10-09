import { useEffect, useState } from "react";
import { useErrorHandler } from "react-error-boundary";

export const useApi = (callBack, initial, handleInErrorBoundry=true) => {
  
  const [state, setState] = useState(initial);
  const handleError = useErrorHandler();

  const fetchData = async () => {
    try {
      const res = await callBack();
      setState(res);
    } catch (error) {
      if (handleInErrorBoundry) {
        handleError(error);
      }
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return {state};
};
