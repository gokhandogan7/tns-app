import { useEffect, useState } from "react";
import { useErrorHandler } from "react-error-boundary";

export const useApi = (callBack, initial, ...args) => {
  const [state, setState] = useState(initial);
  const handleError = useErrorHandler();

  const fetchData = async () => {
    try {
      const res = await callBack(...args);
      setState(res);
    } catch (error) {
        handleError(error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [JSON.stringify(args)]);

  return [state];
};
