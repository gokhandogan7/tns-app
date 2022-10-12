import { useEffect, useState } from "react";
import { useErrorHandler } from "react-error-boundary";

export const useApi = (callBack, initial, searchKey) => {
  console.log(searchKey)
  const [state, setState] = useState(initial);
  const handleError = useErrorHandler();

  const fetchData = async () => {
    try {
      const res = await callBack(searchKey);
      setState(res);
    } catch (error) {
        handleError(error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [searchKey]);

  return {state};
};
