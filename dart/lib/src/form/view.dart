import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:retry/retry.dart';

typedef NextListFunc = Future<int> Function(int);

class InfiniteListView extends StatefulWidget {
  @required
  final NextListFunc nextList;
  final int retryTimeoutSecond;
  final FutureOr<void> Function(Exception) onRetry;
  @required
  final IndexedWidgetBuilder itemBuilder;
  @required
  final int rows;
  @required
  final int fetchWhenRowsLeftToShow;
  @required
  final String noMoreText;
  final double separatorHeight;
  const InfiniteListView({
    Key key,
    this.nextList,
    this.retryTimeoutSecond = 10,
    this.onRetry,
    this.itemBuilder,
    this.rows,
    this.fetchWhenRowsLeftToShow,
    this.noMoreText,
    this.separatorHeight = 0,
  })  : assert(rows >= 20),
        assert(fetchWhenRowsLeftToShow >= 0 && fetchWhenRowsLeftToShow <= rows),
        super(key: key);
  @override
  _InfiniteListViewState createState() => _InfiniteListViewState();
}

class _InfiniteListViewState extends State<InfiniteListView> {
  int page = 0;
  int _total = 0;
  bool _isNoMore = false;
  Future<int> _fetching;

  @override
  void initState() {
    super.initState();
    _nextList();
  }

  @override
  Widget build(BuildContext context) {
    if (widget.separatorHeight > 0)
      return ListView.separated(
        itemCount: _total + 1,
        itemBuilder: _buildItem,
        separatorBuilder: _buildSeparator,
      );
    return ListView.builder(
      itemCount: _total + 1,
      itemBuilder: _buildItem,
    );
  }

  Widget _buildItem(BuildContext context, int idx) {
    // next words
    if (idx >= _total) {
      // at end
      if (_isNoMore)
        return Container(
            alignment: Alignment.center,
            padding: const EdgeInsets.all(16.0),
            child: Text(
              widget.noMoreText,
              style: const TextStyle(color: Colors.grey),
            ));

      if (_total == 0) _nextList();

      // loading
      return Container(
        padding: const EdgeInsets.all(16.0),
        alignment: Alignment.center,
        child: const SizedBox(
            width: 24.0,
            height: 24.0,
            child: CircularProgressIndicator(strokeWidth: 2.0)),
      );
    }

    if (_total - idx == widget.fetchWhenRowsLeftToShow) _nextList();

    // build item
    return widget.itemBuilder(context, idx);
  }

  Widget _buildSeparator(BuildContext context, int index) =>
      Divider(height: widget.separatorHeight);

  void _nextList() {
    if (_isNoMore || _fetching != null) return;
    _fetching = retry(
      () => widget
          .nextList(page)
          .timeout(Duration(seconds: widget.retryTimeoutSecond)),
      retryIf: (e) => e is SocketException || e is TimeoutException,
      onRetry: widget.onRetry,
    );
    _fetching.then((size) {
      if (size == 0) {
        if (!_isNoMore) setState(() => _isNoMore = true);
        return;
      }
      setState(() {
        page++;
        _total += size;
        _isNoMore = size < widget.rows;
      });
    }).whenComplete(() => _fetching = null);
  }
}
