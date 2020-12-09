import pandas as pd


def reverse_handler(event, context):
    """
    Silly approach to reversing an input array that leverages Pandas. This is
    only done as an example of using a third party library with c extensions.
    """
    input_array = event.get("input", [])
    df = pd.DataFrame(input_array)
    return {"output": list(df[::-1][0])}
