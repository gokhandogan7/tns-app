import { useEffect, useState } from "react";
import { useErrorHandler } from "react-error-boundary";

export const useApi = (callBack, initial) => {
  
  const [state, setState] = useState(initial);
  const handleError = useErrorHandler();

  const fetchData = async () => {
    try {
      const res = await callBack();
      setState(res);
    } catch (error) {
        handleError(error);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return {state};
};
