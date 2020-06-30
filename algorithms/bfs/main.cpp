#include <iostream>
#include <list>

using namespace std;

class Graph
{
    int v;          // number of vertiices
    list<int> *adj; // adjency list

public:
    Graph(int v);

    void addEdge(int v, int w);

    void BFS(int s);
};

Graph::Graph(int v)
{
    this->v = v;
    adj = new list<int>[v];
}

void Graph::addEdge(int v, int w)
{
    adj[v].push_back(w);
}

void Graph::BFS(int s)
{
    bool *visited = new bool[v];
    for (int i = 0; i < v; i++)
    {
        visited[i] = false;
    }

    list<int> que;

    visited[s] = true;
    que.push_back(s);

    list<int>::iterator i;

    while (!que.empty())
    {
        s = que.front();
        cout << s << " ";
        que.pop_front();

        for (i = adj[s].begin(); i != adj[s].end(); ++i)
        {
            if (!visited[*i])
            {
                visited[*i] = true;
                que.push_back(*i);
            };
        }
    }
}
