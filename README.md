# Benchmark Results
_See [raw-results.txt][raw-results] for the unformatted output of the
benchmark tests._

[raw-results]: https://git.sr.ht/~smlavine/testing-go-queues/tree/master/item/raw-results.txt

<table>
<tr><th colspan=2>n=10 (1000 iterations)</th></tr>
<tr><td>list</td><td>0m0.012s</td></tr>
<tr><td><b>queue</b></td><td><b>0m0.010s</b></td></tr>
<tr><td><b>slice</b></td><td><b>0m0.010s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=10, random insert (1000 iterations)</th></tr>
<tr><td>list</td><td>0m0.022s</td></tr>
<tr><td><b>queue</b></td><td><b>0m0.020s</b></td></tr>
<tr><td>slice</td><td>0m0.021s</td></tr>
</table>

<table>
<tr><th colspan=2>n=100 (1000 iterations)</th></tr>
<tr><td>list</td><td>0m0.078s</td></tr>
<tr><td>queue</td><td>0m0.076s</td></tr>
<tr><td><b>slice</b></td><td><b>0m0.059s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=100, random insert (1000 iterations)</th></tr>
<tr><td>list</td><td>0m0.199s</td></tr>
<tr><td>queue</td><td>0m0.188s</td></tr>
<tr><td><b>slice</b></td><td><b>0m0.174s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=1000 (1000 iterations)</th></tr>
<tr><td>list</td><td>0m0.813s</td></tr>
<tr><td>queue</td><td>0m1.086s</td></tr>
<tr><td><b>slice</b></td><td><b>0m0.777s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=1000, random insert (1000 iterations)</th></tr>
<tr><td>list</td><td>0m2.572s</td></tr>
<tr><td>queue</td><td>0m2.104s</td></tr>
<tr><td><b>slice</b></td><td><b>0m2.024s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=10000 (1000 iterations)</th></tr>
<tr><td>list</td><td>0m10.676s</td></tr>
<tr><td>queue</td><td>0m11.711s</td></tr>
<tr><td><b>slice</b></td><td><b>0m8.803s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=10000, random insert (1000 iterations)</th></tr>
<tr><td>list</td><td>0m29.539s</td></tr>
<tr><td>queue</td><td>0m24.285s</td></tr>
<tr><td><b>slice</b></td><td><b>0m23.293s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=100000</th></tr>
<tr><td>list</td><td>0m0.074s</td></tr>
<tr><td>queue</td><td>0m0.065s</td></tr>
<tr><td><b>slice</b></td><td><b>0m0.063s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=100000, random insert</th></tr>
<tr><td>list</td><td>0m0.241s</td></tr>
<tr><td><b>queue</b></td><td><b>0m0.205s</b></td></tr>
<tr><td>slice</td><td>0m0.225s</td></tr>
</table>

<table>
<tr><th colspan=2>n=1000000</th></tr>
<tr><td>list</td><td>0m0.844s</td></tr>
<tr><td>queue</td><td>0m0.684s</td></tr>
<tr><td><b>slice</b></td><td><b>0m0.614s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=1000000, random insert</th></tr>
<tr><td>list</td><td>0m2.152s</td></tr>
<tr><td><b>queue</b></td><td><b>0m1.916s</b></td></tr>
<tr><td>slice</td><td>0m2.081s</td></tr>
</table>

<table>
<tr><th colspan=2>n=10000000</th></tr>
<tr><td>list</td><td>0m7.843s</td></tr>
<tr><td>queue</td><td>0m6.955s</td></tr>
<tr><td><b>slice</b></td><td><b>0m6.082s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=10000000, random insert</th></tr>
<tr><td>list</td><td>0m20.915s</td></tr>
<tr><td>queue</td><td>0m19.353s</td></tr>
<tr><td><b>slice</b></td><td><b>0m19.030s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=100000000</th></tr>
<tr><td>list</td><td>1m16.619s</td></tr>
<tr><td>queue</td><td>1m6.998s</td></tr>
<tr><td><b>slice</b></td><td><b>1m1.921s</b></td></tr>
</table>

<table>
<tr><th colspan=2>n=100000000, random insert</th></tr>
<tr><td>list</td><td>3m42.505s</td></tr>
<tr><td>queue</td><td>3m13.185s</td></tr>
<tr><td><b>slice</b></td><td><b>3m6.901s</b></td></tr>
</table>
