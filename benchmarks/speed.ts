import { performance } from "perf_hooks";

export const benchmark = (fn: () => unknown) => {
  const start = performance.now();
  const result = fn();
  console.log(
    "Execution concluded in:",
    (performance.now() - start).toFixed(5),
    "ms"
  );
  console.log("Execution result:", result);
};
