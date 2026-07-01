import re

def vertical_histogram_of(s: str) -> str:

    if len(s) == 0:
        return ""
    
    s = re.sub("[^A-Z]", "", s)
    
    vals = sorted(set(s))
    l = {v: 0 for v in vals}

    for i in s:
        l[i] += 1
    
    if len(s) > 0:
        max_l = max(l.values())

    str_v = ""
    for i in range(max_l, 0, -1):
        for key,val in l.items():
            if val >= i:
                str_v += "* "
            else:
                str_v += "  "
        str_v = str_v.rstrip(" ")
        str_v += "\n"
    
    str_v += ' '.join(vals)
    
    return str_v