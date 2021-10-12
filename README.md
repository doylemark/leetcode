## Typescript
* To run a typescript file run `npm run file --file=somefile.ts`

* For a more brief version (`lc somefile.ts`) a zsh alias will suffice:
```zsh
function lc() {
  npm run file --file=$1
}
```