module.exports = {
  mode: "jit",
  purge: ["./**/*.go"],
  darkMode: "class", // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        primary: "var(--fg)",
        red: {
          DEFAULT: "var(--red)",
          bg: "var(--bg_red)",
          diff: "var(--diff_red)",
        },
        green: {
          DEFAULT: "var(--green)",
          bg: "var(--bg_green)",
          diff: "var(--diff_green)",
        },
        blue: {
          DEFAULT: "var(--blue)",
          bg: "var(--bg_blue)",
          diff: "var(--diff_blue)",
        },
        orange: "var(--orange)",
        yellow: "var(--yellow)",
        purple: "var(--purple)",
        grey: "var(--grey)",
        background: {
          black: "var(--black)",
          DEFAULT: "var(--bg0)",
        },
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
