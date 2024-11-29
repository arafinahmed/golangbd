ReactDOM.render(
    React.createElement(GraphiQL, {
      fetcher: GraphiQL.createFetcher({
        url: window.location.origin + '/graphql',
      }),
      defaultEditorToolsVisibility: true,
      query:  "# Welcome to BOSEMANN GraphQL Playground\n"+
              "#\n"+
              "# GraphiQL is an in-browser tool for writing, validating, and\n"+
              "# testing GraphQL queries.\n"+
              "#\n"+
              "# Type queries into this side of the screen, and you will see intelligent\n"+
              "# typeaheads aware of the current GraphQL type schema and live syntax and\n"+
              "# validation errors highlighted within the text.\n"+
              "#\n"+
              "# GraphQL queries typically start with a \"{\" character. Lines that start\n"+
              "# with a # are ignored.\n"+
              "#\n"+
              "# An example GraphQL query might look like:\n"+
              "#\n"+
              "#     {\n"+
              "#       listPackage{\n"+
              "#         id\n"+
              "#         name\n"+   
              "#       }\n"+
              "#     }\n"+
              "#\n"+
              "# Keyboard shortcuts:\n"+
              "#\n"+
              "#   Prettify query:  Shift-Ctrl-P (or press the prettify button)\n"+
              "#\n"+
              "#  Merge fragments:  Shift-Ctrl-M (or press the merge button)\n"+
              "#\n"+
              "#        Run Query:  Ctrl-Enter (or press the play button)\n"+
              "#\n"+
              "#    Auto Complete:  Ctrl-Space (or just start typing)\n"+
              "#\n",
            //theme: "light",
      //variables: 'ss',
      //defaultQuery:"test",
    }),
    document.getElementById('graphiql'),
  );

